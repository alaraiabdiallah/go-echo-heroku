package main

import (
	"os"

	"github.com/alaraiabdiallah/golang-heroku-starter/src/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := os.Getenv("PORT")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	routes.APIV1(api)

	e.Logger.Fatal(e.Start(":" + port))
}
