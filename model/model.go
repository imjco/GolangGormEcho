package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	LastName     string     `json:"last_name"`
	FirstName    string     `json:"first_name"`
	MiddleName   string     `json:"middle_name"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	Address      string     `json:"address"`
	DateofBirth  string     `json:"date_of_birth"`
	GenderTypeID uint       `json:"gender_type_id"`
	GenderType   GenderType `json:"gender_type"`
}

type GenderType struct {
	gorm.Model
	GenderName string `json:"name" gorm:"unique;size:50"`
}
