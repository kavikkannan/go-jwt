package models

type Login struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Position string
	Password []byte `json:"-"`
	Phonenumber string `json:"phonenumber"`
	Address string `json:"address"`
}