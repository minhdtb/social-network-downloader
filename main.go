package main

import (
	"./plugins"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientRequest struct {
	Plugin int32  `json:"plugin"`
	Url    string `json:"url"`
}

type ClientResponse struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	VideoUrl  string `json:"videoUrl"`
}

func main() {
	e := echo.New()

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.POST("/get-content", func(context echo.Context) error {
		clientRequest := new(ClientRequest)

		_ = context.Bind(clientRequest)

		var netClient = http.Client{
			Timeout: time.Second * 10,
		}

		response, _ := netClient.Get(clientRequest.Url)

		c, _ := ioutil.ReadAll(response.Body)

		content := string(c)

		var plugin plugins.Plugin
		if clientRequest.Plugin == 0 {
			plugin = plugins.Facebook{}
		} else {
			plugin = plugins.Instagram{}
		}

		clientResponse := new(ClientResponse)

		var title = plugin.GetTitle(content)
		var thumbnail = plugin.GetThumbnail(content)
		var videoUrl = plugin.GetVideoUrl(content)

		if title != nil {
			clientResponse.Title = *title
		} else {
			clientResponse.Title = ""
		}

		if thumbnail != nil {
			clientResponse.Thumbnail = *thumbnail
		} else {
			clientResponse.Thumbnail = ""
		}

		if videoUrl != nil {
			clientResponse.VideoUrl = *videoUrl
		} else {
			clientResponse.VideoUrl = ""
		}

		return context.JSON(http.StatusOK, clientResponse)
	})

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Works!!!!!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
