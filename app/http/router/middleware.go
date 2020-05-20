package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)

// Обрабатывает панику и пишет в логи
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				request, _ := httputil.DumpRequest(c.Request, false)

				logrus.WithField("request", string(request)).
					Errorf("Перехвачена паника: %s", err)
				c.AbortWithError(http.StatusInternalServerError, errors.New("Critical error during request"))
			}
		}()

		c.Next()
	}
}
