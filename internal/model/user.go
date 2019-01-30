package model

type User struct {
	Id      int `gorm:"primary_key"`
	name    string
	email   string
	picture string
	Audit
}
