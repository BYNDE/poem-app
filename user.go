package platform

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
