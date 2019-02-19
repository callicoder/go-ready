package model

type Group struct {
	Id          uint64 `gorm:"primary_key"`
	Name        string
	Description string
	ImageUrl    string
	CreatedBy   uint64
	UpdatedBy   uint64
	Audit
}
