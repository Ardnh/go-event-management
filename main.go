package main

import (
	"go/ems/app"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.DbConnection()
	validate := validator.New()

	router := app.Router(db, validate)

	router.Run(":8080")
}
