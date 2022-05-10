package controller

import (
	"main/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEndPoint(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, twippy api")
}

func GetTweets(c echo.Context) error {
	name := c.QueryParam("name")

	tweets := model.FetchTweet(name)
	return c.JSON(http.StatusOK, tweets)
}
