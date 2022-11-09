package middleware

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
	"github.com/google/uuid"
	"time"
)

func AccessLogMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			requestProcessStartTime := time.Now()

			defer func() {
				logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("Request process finished. Spent time: ", time.Since(requestProcessStartTime))
			}()
			//c.Request().Header.Set(echo.HeaderXRequestID, uuid.New().String())
			ctx.SetRequest(ctx.Request().Clone(context.WithValue(ctx.Request().Context(), domain.REQUEST_ID_KEY_FOR_CONTEXT, uuid.New().String())))

			r := ctx.Request()
			logger.LogrusLoggerWithContext(ctx.Request().Context()).Info("Request method: ", r.Method, ", remote address: ", r.RemoteAddr, ", request URL: ", r.URL.Path, ", request process start time: ", requestProcessStartTime)

			//csrf := r.Header.Get("X-XSRF-Token")
			//k, _ := r.Cookie("_csrf")
			//assert.Equal(k, csrf)
			return next(ctx)
		}
	}
}
