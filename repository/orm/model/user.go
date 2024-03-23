package model

type User struct {
	ID       int    `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func (User) TableName() string {
	return "user"
}
