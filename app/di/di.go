package di

import (
	"github.com/gin-gonic/gin"
	"otus/messages-counter/app/usecase"
	"otus/messages-counter/provider"
)

// Инстанс приложения
type DI struct {
	CacheProvider provider.Cachier
	MessagesAPI   usecase.MessagesAPI
}

// Новый инстанс приложения
func NewDI(cacheProvider provider.Cachier, messagesAPI usecase.MessagesAPI) DI {
	return DI{
		CacheProvider: cacheProvider,
		MessagesAPI:   messagesAPI,
	}
}

// Пробрасывает зависимости в хэндлеры
func (di DI) ProvideDependency(f func(c *gin.Context, di DI)) func(*gin.Context) {
	return func(c *gin.Context) {
		f(c, di)
	}
}
