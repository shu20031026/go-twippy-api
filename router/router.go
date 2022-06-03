package router

import (
	"fmt"
	"main/controller"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".envファイルの読み込みに失敗しました: %v", err)
	}
	var PORT string
	if os.Getenv("PORT") == "" {
		PORT = "80"
	} else {
		PORT = os.Getenv("PORT")
	}
	e := echo.New()

	e.GET("/", controller.GetEndPoint)
	e.GET("/tweets", controller.GetTweets)

	e.Logger.Fatal(e.Start(":" + PORT))
}
