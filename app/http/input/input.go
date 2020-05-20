package input

// Устанавливаю счетчик
type SetNbUnread struct {
	UserId int `form:"userId" binding:"required"`
	Cnt    int `form:"cnt" binding:"required"`
}

// Получаю счетчик
type GetNbUnread struct {
	UserId int `form:"userId" binding:"required"`
}
