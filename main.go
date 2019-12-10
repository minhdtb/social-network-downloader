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
	plugin int32
	url    string
}

type ClientResponse struct {
	title     string
	thumbnail string
	videoUrl  string
}

func main() {
	e := echo.New()

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

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

		response, _ := netClient.Get(clientRequest.url)

		c, _ := ioutil.ReadAll(response.Body)

		content := string(c)

		var plugin plugins.Plugin
		if clientRequest.plugin == 0 {
			plugin = plugins.Facebook{}
		} else {
			plugin = plugins.Instagram{}
		}

		clientResponse := new(ClientResponse)

		var title = plugin.GetTitle(content)
		var thumbnail = plugin.GetThumbnail(content)
		var videoUrl = plugin.GetVideoUrl(content)

		if title != nil {
			clientResponse.title = *title
		} else {
			clientResponse.title = ""
		}

		if thumbnail != nil {
			clientResponse.thumbnail = *thumbnail
		} else {
			clientResponse.thumbnail = ""
		}

		if videoUrl != nil {
			clientResponse.videoUrl = *videoUrl
		} else {
			clientResponse.videoUrl = ""
		}

		return context.JSON(http.StatusOK, clientResponse)
	})

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Works!!!!!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
