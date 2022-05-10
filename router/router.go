package router

import (
	"main/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controller.EndPoint)

	e.Logger.Fatal(e.Start(":80"))
}
