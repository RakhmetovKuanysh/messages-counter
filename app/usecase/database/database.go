package database

import (
	"otus/messages-counter/app/usecase"
)

// Хранилище счетчиков
type CounterDatabase struct {
}

// Новое хранилище счетчиков
func NewCounterDatabase() usecase.CounterDatabase {
	return &CounterDatabase{}
}
