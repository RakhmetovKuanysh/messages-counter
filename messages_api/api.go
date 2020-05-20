package messages_api

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"otus/messages-counter/app/usecase"
	"strconv"
)

// API
type MessagesAPI struct {
}

// Новый клиент
func NewMessagesAPI() usecase.MessagesAPI {
	return &MessagesAPI{}
}

// Ответ от API
type GetNbUnreadResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cnt     int    `json:"cnt"`
}

// Получает количество непрочитанных сообщений
func (r *MessagesAPI) GetNbUnread(userId int) (cnt int, err error) {
	resp, err := http.Get("http://127.0.0.1:8080/get-nb-unread?userId=" + strconv.Itoa(userId))

	if err != nil {
		logrus.WithError(err).Error("Ошибка при простроении запроса для получения количества непрочитанных сообщений")
		return
	}

	defer resp.Body.Close()

	var result GetNbUnreadResult
	json.NewDecoder(resp.Body).Decode(&result)

	if result.Code != http.StatusOK {
		return 0, errors.New("ошибка при получении количества сообщений")
	}

	cnt = result.Cnt

	return
}
