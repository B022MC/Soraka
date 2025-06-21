package ginApp

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"Soraka/pkg/fastcurd"
	"Soraka/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// ctx key
const (
	KeyApp           = "app"
	KeyInitOnce      = "initOnce"
	KeyProcBeginTime = "procBeginTime"
	KeyProcEndTime   = "procEndTime"

	KeyResp       = "resp"
	KeyReqID      = "reqID"
	KeyStatusCode = "statusCode"
	KeyProcTime   = "procTime"
)

var (
	respServerBad    = &fastcurd.RetJSON{Code: fastcurd.CodeServerError, Msg: "服务器开小差了~"}
	respBadReq       = &fastcurd.RetJSON{Code: fastcurd.CodeBadReq, Msg: "错误的请求"}
	respNoChange     = &fastcurd.RetJSON{Code: fastcurd.CodeDefaultError, Msg: "无更新"}
	respNoAuth       = &fastcurd.RetJSON{Code: fastcurd.CodeNoAuth, Msg: "未授权"}
	respNoLogin      = &fastcurd.RetJSON{Code: fastcurd.CodeNoLogin, Msg: "未登录"}
	respReqFrequency = &fastcurd.RetJSON{Code: fastcurd.CodeRateLimitError, Msg: "请求速度太快了~"}
)

type (
	App struct {
		C    *gin.Context
		mu   sync.Mutex
		sqls []logger.SqlRecord
	}
)

func NewGin() *gin.Engine {
	engine := gin.New()
	engine.Use(PrepareProc)
	return engine
}

func GetApp(c *gin.Context) *App {
	initOnce, ok := c.Get(KeyInitOnce)
	if !ok {
		panic("ctx must set initOnce")
	}
	initOnce.(*sync.Once).Do(func() {
		c.Set(KeyApp, newApp(c))
	})
	app, _ := c.Get(KeyApp)
	return app.(*App)
}
func newApp(c *gin.Context) *App {
	app := &App{
		C:    c,
		sqls: make([]logger.SqlRecord, 0, 4),
	}
	app.setRecordSqlFn()
	return app
}

// finally resp fn

func (app *App) Response(code int, json *fastcurd.RetJSON) {
	reqID := app.GetReqID()
	procBeginTime := app.GetProcBeginTime()
	procEndTime := time.Now()
	procTime := procEndTime.Sub(*procBeginTime)
	json.Extra = &fastcurd.RespJsonExtra{
		ProcTime: procTime.String(),
		ReqID:    reqID,
	}
	if gin.IsDebugging() {
		json.Extra.SQLs = app.sqls
	}
	app.SetProcEndTime(&procEndTime)
	app.SetProcTime(procTime)
	app.SetCtxRespVal(json)
	app.SetStatusCode(code)
	app.C.JSON(code, json)
	app.C.Abort()
}

// resp helper

func (app *App) Ok(msg string, params ...any) {
	var actData any = nil
	if len(params) == 1 {
		actData = params[0]
	}
	app.JSON(&fastcurd.RetJSON{Code: fastcurd.CodeOk, Msg: msg, Data: actData})
}
func (app *App) Data(data any) {
	app.RetData(data)
}
func (app *App) ServerError(err error) {
	json := &fastcurd.RetJSON{Code: fastcurd.CodeServerError, Msg: err.Error()}
	app.Response(http.StatusOK, json)
}
func (app *App) ServerBad() {
	app.Response(http.StatusOK, respServerBad)
}
func (app *App) RetData(data any, msgParam ...string) {
	msg := ""
	if len(msgParam) == 1 {
		msg = msgParam[0]
	}
	app.Ok(msg, data)
}
func (app *App) JSON(json *fastcurd.RetJSON) {
	app.Response(http.StatusOK, json)
}
func (app *App) SendList(list any, count int) {
	app.Response(http.StatusOK, &fastcurd.RetJSON{
		Code:  fastcurd.CodeOk,
		Data:  list,
		Count: &count,
	})
}
func (app *App) BadReq() {
	app.Response(http.StatusOK, respBadReq)
}
func (app *App) String(html string) {
	app.C.String(http.StatusOK, html)
}
func (app *App) ValidError(err error) {
	json := &fastcurd.RetJSON{}
	var actErr validator.ValidationErrors
	switch {
	case errors.As(err, &actErr):
		json.Code = fastcurd.CodeValidError
		json.Msg = actErr[0].Error()
	default:
		if err.Error() == "EOF" {
			json.Code = fastcurd.CodeValidError
			json.Msg = "request param is required"
		} else {
			json.Code = fastcurd.CodeDefaultError
			json.Msg = actErr.Error()
		}
	}
	app.Response(http.StatusOK, json)
}
func (app *App) NoChange() {
	app.JSON(respNoChange)
}
func (app *App) NoAuth() {
	app.Response(http.StatusUnauthorized, respNoAuth)
}
func (app *App) NoLogin() {
	app.Response(http.StatusUnauthorized, respNoLogin)
}
func (app *App) ErrorMsg(msg string) {
	json := &fastcurd.RetJSON{Code: fastcurd.CodeDefaultError, Msg: msg}
	app.Response(http.StatusOK, json)
}
func (app *App) CommonError(err error) {
	app.ErrorMsg(err.Error())
}
func (app *App) RateLimitError() {
	app.Response(http.StatusOK, respReqFrequency)
}
func (app *App) Success() {
	json := &fastcurd.RetJSON{Code: fastcurd.CodeOk}
	app.Response(http.StatusOK, json)
}
func (app *App) SendAffectRows(affectRows int) {
	app.Data(gin.H{
		"affectRows": affectRows,
	})
}

// sql

func (app *App) RecordSql(record logger.SqlRecord) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.sqls = append(app.sqls, record)
}
func (app *App) GetSqls() []logger.SqlRecord {
	return app.sqls
}

// ctx value helper
func (app *App) setRecordSqlFn() {
	ctx := app.C.Request.Context()
	ctx = logger.SetRecordSqlFn(ctx, app.RecordSql)
	app.C.Request = app.C.Request.WithContext(ctx)
}

func (app *App) GetCtxRespVal() *fastcurd.RetJSON {
	if json, ok := app.C.Get(KeyResp); ok {
		return json.(*fastcurd.RetJSON)
	}
	return nil
}
func (app *App) SetCtxRespVal(json *fastcurd.RetJSON) {
	app.C.Set(KeyResp, json)
}

func (app *App) GetProcBeginTime() *time.Time {
	procBeginTime, _ := app.C.Get(KeyProcBeginTime)
	return procBeginTime.(*time.Time)
}

func (app *App) GetProcEndTime() *time.Time {
	if procBeginTime, ok := app.C.Get(KeyProcBeginTime); ok {
		return procBeginTime.(*time.Time)
	}
	return nil
}
func (app *App) SetProcEndTime(t *time.Time) {
	app.C.Set(KeyProcEndTime, t)
}
func (app *App) GetReqID() string {
	reqID, _ := app.C.Get(KeyReqID)
	return reqID.(string)
}
func (app *App) GetStatusCode() int {
	if code, ok := app.C.Get(KeyStatusCode); ok {
		return code.(int)
	}
	return 0
}
func (app *App) SetStatusCode(code int) {
	app.C.Set(KeyStatusCode, code)
}
func (app *App) GetProcTime() time.Duration {
	if procTime, ok := app.C.Get(KeyProcTime); ok {
		return procTime.(time.Duration)
	}
	return 0
}
func (app *App) SetProcTime(t time.Duration) {
	app.C.Set(KeyProcTime, t)
}

// middleware

func PrepareProc(c *gin.Context) {
	begin := time.Now()
	c.Set(KeyInitOnce, &sync.Once{})
	c.Set(KeyProcBeginTime, &begin)   // 请求开始时间
	c.Set(KeyReqID, uuid.NewString()) // 请求id
	c.Next()
}
