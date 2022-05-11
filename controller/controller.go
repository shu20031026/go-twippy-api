package controller

import (
	"main/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEndPoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, twippy api")
}

// 型定義
type tweetDataType struct {
	Name       string   `json:"name"`
	ScreenName string   `json:"screenName"`
	Icon       string   `json:"icon"`
	Tweets     []string `json:"tweets"`
}

func GetTweets(c echo.Context) error {
	name := c.QueryParam("name")

	tweets := model.FetchTweet(name)
	return c.JSON(http.StatusOK, tweets)
}
