package response

import (
	"Soraka/pkg/fastcurd"
	"Soraka/pkg/logger"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Response(c *gin.Context, code int, json *fastcurd.RetJSON, reqID string, procBegin time.Time, sqls []logger.SqlRecord) {
	procEnd := time.Now()
	procTime := procEnd.Sub(procBegin)

	json.Extra = &fastcurd.RespJsonExtra{
		ProcTime: procTime.String(),
		ReqID:    reqID,
	}
	if gin.IsDebugging() {
		json.Extra.SQLs = sqls
	}

	c.Set("statusCode", code)
	c.Set("procEndTime", procEnd)
	c.Set("procTime", procTime)
	c.Set("resp", json)

	c.JSON(code, json)
	c.Abort()
}

func Success(c *gin.Context) {
	Response(c, http.StatusOK, &fastcurd.RetJSON{Code: fastcurd.CodeOk}, getReqID(c), getProcBegin(c), nil)
}

func Ok(c *gin.Context, msg string, data any) {
	Response(c, http.StatusOK, &fastcurd.RetJSON{
		Code: fastcurd.CodeOk,
		Msg:  msg,
		Data: data,
	}, getReqID(c), getProcBegin(c), nil)
}

func Data(c *gin.Context, data any) {
	Ok(c, "", data)
}

func ErrorMsg(c *gin.Context, msg string) {
	Response(c, http.StatusOK, &fastcurd.RetJSON{Code: fastcurd.CodeDefaultError, Msg: msg}, getReqID(c), getProcBegin(c), nil)
}

func CommonError(c *gin.Context, err error) {
	ErrorMsg(c, err.Error())
}

func ValidError(c *gin.Context, err error) {
	var actErr validator.ValidationErrors
	json := &fastcurd.RetJSON{}
	if errors.As(err, &actErr) {
		// 是验证错误
		json.Code = fastcurd.CodeValidError
		json.Msg = actErr[0].Error()
	} else if err.Error() == "EOF" {
		json.Code = fastcurd.CodeValidError
		json.Msg = "request param is required"
	} else {
		json.Code = fastcurd.CodeDefaultError
		json.Msg = err.Error()
	}
	Response(c, http.StatusOK, json, getReqID(c), getProcBegin(c), nil)
}

func SendList(c *gin.Context, list any, count int) {
	Response(c, http.StatusOK, &fastcurd.RetJSON{
		Code:  fastcurd.CodeOk,
		Data:  list,
		Count: &count,
	}, getReqID(c), getProcBegin(c), nil)
}

func getReqID(c *gin.Context) string {
	reqID, _ := c.Get("reqID")
	return reqID.(string)
}

func getProcBegin(c *gin.Context) time.Time {
	t, _ := c.Get("procBeginTime")
	return *(t.(*time.Time))
}
