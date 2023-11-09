package main

import (
	"profile/db"
	"profile/model"
)

func main() {
	db.InitializeDatabases()
	err := db.GormDB["default"].AutoMigrate(
		&model.Customer{},
		&model.GenderType{},
	)
	if err != nil {
		panic(err)
	}
}
