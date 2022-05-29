package router

import (
	"main/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Printf(".envファイルの読み込みに失敗しました: %v", err)
	// }
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controller.GetEndPoint)
	e.GET("/tweets", controller.GetTweets)

	e.Start(":" + os.Getenv("PORT"))
	// e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
