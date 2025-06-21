package middleware

import (
	"errors"

	ginApp "Soraka/pkg/gin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryWithLogFn(logFn func(msg string, keysAndVals ...any)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			app := ginApp.GetApp(c)
			reqID := app.GetReqID()
			var actErr error
			logFn("bdk.gin.recover", zap.Any("recoverErr", err), zap.String("reqID", reqID),
				zap.StackSkip("errStack", 2))
			switch err := err.(type) {
			case error:
				actErr = err
			case string:
				errMsg := err
				actErr = errors.New(errMsg)
			default:
				actErr = errors.New("server panic")
			}
			app.ServerError(actErr)
		}()
		c.Next()
	}
}
