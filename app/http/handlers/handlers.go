package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"otus/messages-counter/app/di"
	app "otus/messages-counter/app/http"
	"otus/messages-counter/app/http/input"
	"strconv"
)

// Проверка состояния сервиса
func Health(c *gin.Context, di di.DI) {
	c.String(http.StatusOK, "ok")
}

// Устанавливает счетчик непрочитанных сообщений
func SetNbUnread(c *gin.Context, di di.DI) {
	in := input.SetNbUnread{}

	if err := c.MustBindWith(&in, binding.Form); err != nil {
		return
	}

	if in.UserId == 0 {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide parameters"))

		return
	}

	di.CacheProvider.Set(strconv.Itoa(in.UserId), in.Cnt, 86400)

	c.JSON(http.StatusOK, app.WithSuccess("Updated"))

	return
}

// Получает счетчик непрочитанных сообщений
func GetNbUnread(c *gin.Context, di di.DI) {
	userId, ok := c.GetQuery("userId")

	if !ok {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide parameters"))

		return
	}

	receiverid, err := strconv.Atoi(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide valid parameters"))
		return
	}

	nbCnt, err := di.CacheProvider.GetInt(userId)

	if err != nil {
		nbCnt, err = di.MessagesAPI.GetNbUnread(receiverid)
	}

	c.JSON(http.StatusOK, app.GetNbUnreadResponse{
		Response: app.WithSuccess("Found"),
		Cnt:      nbCnt,
	})
}

// Сброс количества сообщений
func UnsetNbUnread(c *gin.Context, di di.DI) {
	userId, ok := c.GetQuery("userId")

	if !ok {
		c.JSON(http.StatusBadRequest, app.WithError(app.PARAMETERS_REQUIRED, "Provide parameters"))

		return
	}

	di.CacheProvider.Delete(userId)

	c.JSON(http.StatusOK, app.WithSuccess("Success"))
}
