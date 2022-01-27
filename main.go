package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var consumerKey = ""
var consumerSecret = ""
var accessToken = ""
var accessTokenSecret = ""

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	timelineQuery := url.Values{}

	timelineQuery.Set("screen_name", "fukke0906")
	timelineQuery.Set("count", "20")
	timelineQuery.Set("include_rts", "false")
	fmt.Println(timelineQuery)
	tweets, err := api.GetUserTimeline(timelineQuery)
	if err != nil {
		fmt.Printf("Error to getHomeTimeline. err:%v\n", err)
		os.Exit(1)
	}

	fmt.Println(tweets)
	e.GET("/", hello)
	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":80"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, twippy-api")
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
