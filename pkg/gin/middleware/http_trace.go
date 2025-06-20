package middleware

import (
	"Soraka/pkg"
	"Soraka/pkg/fastcurd"
	"bytes"

	"io"
	"strings"

	"github.com/gin-gonic/gin"

	ginApp "Soraka/pkg/gin"
	"go.uber.org/zap"
)

const (
	MaxTraceSize   = 1 << 10
	KeySaveResp    = "bdk.saveResp"
	KeyShouldTrace = "bdk.shouldTrace"
)

const (
	ContentTypeJson = "application/json"
)
const (
	shouldTraceNone  = 0
	shouldTraceTrue  = 1
	shouldTraceFalse = 2
)

func NewHttpTrace(logFn func(msg string, keysAndVals ...any)) func(c *gin.Context) {
	return NewHttpTraceWithDefaultTraceParam(logFn, true)
}
func NewHttpTraceWithDefaultNotTrace(logFn func(msg string, keysAndVals ...any)) func(c *gin.Context) {
	return NewHttpTraceWithDefaultTraceParam(logFn, false)
}
func NewHttpTraceWithDefaultTraceParam(logFn func(msg string, keysAndVals ...any), isDefaultTrace bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		statusCode := c.Writer.Status()
		isShouldTraceVal := isShouldTrace(c)
		if (isDefaultTrace && isShouldTraceVal == shouldTraceFalse) ||
			(!isDefaultTrace && isShouldTraceVal != shouldTraceTrue) ||
			pkg.IsSkipLogReq(c.Request, statusCode) {
			return
		}
		app := ginApp.GetApp(c)
		reqID := app.GetReqID()
		reqTime := app.GetProcBeginTime()
		postData := ""
		var resp *fastcurd.RetJSON
		contentType := strings.Split(pkg.GetContentType(c.Request), ";")[0]
		switch contentType {
		case ContentTypeJson:
			bodyBts, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBts))
			postData = pkg.Bytes2Str(bodyBts)
		}
		procTime := app.GetProcTime()
		if isShouldSaveResp(c) {
			resp = app.GetCtxRespVal()
		}
		if len(postData) > MaxTraceSize {
			postData = postData[:MaxTraceSize]
		}
		keysAndValues := []any{
			zap.String("bdk.http.url", c.Request.RequestURI),
			zap.String("bdk.http.clientIP", c.ClientIP()),
			zap.Int("bdk.http.statusCode", statusCode),
			zap.String("bdk.gin.reqID", reqID),
			zap.Duration("bdk.gin.procTime", procTime),
			zap.Timep("bdk.gin.reqTime", reqTime),
		}
		if resp != nil {
			keysAndValues = append(keysAndValues, zap.Any("bdk.http.resp", resp))
		}
		if postData != "" {
			keysAndValues = append(keysAndValues, zap.String("bdk.gin.postData", postData))
		}
		logFn("bdk.httpTrace", keysAndValues...)
	}
}

func isShouldSaveResp(c *gin.Context) bool {
	b, ok := c.Get(KeySaveResp)
	if !ok {
		return false
	}
	if t, ok := b.(bool); ok {
		return t
	}
	return false
}
func setSaveResp(c *gin.Context) {
	c.Set(KeySaveResp, true)
}
func isShouldTrace(c *gin.Context) int {
	b, ok := c.Get(KeyShouldTrace)
	if !ok {
		return shouldTraceNone
	}
	if t, ok := b.(int); ok {
		return t
	}
	return shouldTraceNone
}
func setTrace(c *gin.Context) {
	setShouldTraceVal(c, shouldTraceTrue)
}
func setNotTrace(c *gin.Context) {
	setShouldTraceVal(c, shouldTraceFalse)
}
func setShouldTraceVal(c *gin.Context, val int) {
	c.Set(KeyShouldTrace, val)
}
func SaveResp(c *gin.Context) {
	setSaveResp(c)
	c.Next()
}
func NotTrace(c *gin.Context) {
	setNotTrace(c)
	c.Next()
}
func Trace(c *gin.Context) {
	setTrace(c)
	c.Next()
}
