package router

import (
	"github.com/gin-gonic/gin"
	"otus/messages-counter/app/di"
	"otus/messages-counter/app/http/handlers"
)

// Инициализирует маршрутизатор
func Router(p *di.DI, debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if debug {
		r.Use(gin.Logger())
	}

	r.Use(Recovery())

	configureRoutes(r, p)

	return r
}

// Настройки маршрутов
func configureRoutes(r *gin.Engine, p *di.DI) {
	r.GET("/health", p.ProvideDependency(handlers.Health))
	r.POST("/set-nb-unread", p.ProvideDependency(handlers.SetNbUnread))
	r.GET("/get-nb-unread", p.ProvideDependency(handlers.GetNbUnread))
	r.GET("/unset-nb-unread", p.ProvideDependency(handlers.UnsetNbUnread))
}
