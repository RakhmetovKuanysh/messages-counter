package usecase

// Клиент для работы с MessagesAPI
type MessagesAPI interface {
	GetNbUnread(userId int) (cnt int, err error)
}
