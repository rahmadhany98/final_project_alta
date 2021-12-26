package middleware

import (
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		start := time.Now()

		res := next(c)

		logrus.WithFields(logrus.Fields{
			"method":  c.Request().Method,
			"path":    c.Path(),
			"status":  c.Response().Status,
			"latency": time.Since(start).Milliseconds(),
		}).Info("Request Details")
		return res
	}
}
