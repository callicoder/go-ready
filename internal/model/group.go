package model

type Group struct {
	Id   int `gorm:"primary_key"`
	Name string
	Audit
}
