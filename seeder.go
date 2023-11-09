package main

import (
	"profile/db"
	"profile/model"
)

func main() {
	db.InitializeDatabases()

	var gendertype []model.GenderType
	gendertype = append(gendertype, model.GenderType{
		GenderName: "Male",
	})
	gendertype = append(gendertype, model.GenderType{
		GenderName: "Female",
	})
	res := db.GormDB["default"].Create(&gendertype)
	if res.Error != nil {
		panic(res.Error)
	} else {
		println("GenderType Seeded")
	}

	var customer []model.Customer
	customer = append(customer, model.Customer{
		LastName:     "Doe",
		FirstName:    "John",
		MiddleName:   "Smith",
		Email:        "sample@gmail.com",
		Phone:        "09123456789",
		Address:      "Sample Address",
		DateofBirth:  "1992-01-01",
		GenderTypeID: 1,
	})
	customer = append(customer, model.Customer{
		LastName:     "Doe",
		FirstName:    "Jane",
		MiddleName:   "Smith",
		Email:        "sample2@gmail.com",
		Phone:        "09584637291",
		Address:      "Sample Address",
		DateofBirth:  "1990-02-01",
		GenderTypeID: 2,
	})

	res = db.GormDB["default"].Create(&customer)
	if res.Error != nil {
		panic(res.Error)
	} else {
		println("Customer Seeded")
	}

}
