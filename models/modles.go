package models

type Model struct {
	ID         int `gorm:"primary_key"`
	CreatedOn  int
	ModifiedOn int
}
