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
type MoldDataType struct {
	Name       string   `json:"name"`
	ScreenName string   `json:"screenName"`
	Icon       string   `json:"icon"`
	Tweets     []string `json:"tweets"`
}

func GetTweets(c echo.Context) error {
	name := c.QueryParam("name")
	// count := c.QueryParam("count")

	tweets := model.FetchTweet(name)

	var moldData MoldDataType
	moldData.Icon = tweets[0].User.ProfileImageUrlHttps
	moldData.Name = tweets[0].User.Name
	moldData.ScreenName = tweets[0].User.ScreenName
	var TweetsSlice = moldData.Tweets

	for _, tweet := range tweets {
		TweetsSlice = append(TweetsSlice, tweet.FullText)
	}
	moldData.Tweets = TweetsSlice

	return c.JSON(http.StatusOK, moldData)
}
