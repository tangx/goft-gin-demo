package models

type User struct {
	ID        int    `gorm:""`
	Name      string `gorm:""`
	Age       int    `gorm:""`
	Cellphone string `gorm:""`
}
