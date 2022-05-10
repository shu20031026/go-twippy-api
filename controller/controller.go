package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var consumerKey string
var consumerSecret string
var accessToken string
var accessTokenSecret string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".envファイルの読み込みに失敗しました: %v", err)
	}
	consumerKey = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessToken = os.Getenv("ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
}

func EndPoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, twippy api")
}
