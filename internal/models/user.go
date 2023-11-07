package models

import "gorm.io/gorm"

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
}

type UserApplication struct {
	Name          string `json:"name"`
	InstituteName string `json:"institute_name"`
	Jid           uint   `json:"jid"`
	Jobs          Jobs   `json:"job_application"`
}
