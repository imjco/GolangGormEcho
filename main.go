package main

import (
	"github.com/labstack/echo/v4"
	"profile/api"
)

func main() {
	e := echo.New()

	api.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))

}
