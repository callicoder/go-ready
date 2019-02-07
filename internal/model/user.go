package model

type User struct {
	Id       int `gorm:"primary_key"`
	Name     string
	Email    string
	ImageUrl string
	Audit
}
