package models

type User struct {
	Model
	Username   string
	Password   string
	CreatedBy  string
	ModifiedBy string
	State      int
}
