package model

type User struct {
	Id        uint64 `gorm:"primary_key"`
	FirstName string
	LastName  string
	Email     string
	ImageUrl  string
	Audit
}
