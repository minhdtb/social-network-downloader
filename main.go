package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/minhdtb/social-network-downloader/plugins"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type ClientRequest struct {
	Url string `json:"url"`
}

type ClientResponse struct {
	Url         string `json:"url"`
	Type        int32  `json:"type"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	VideoUrl    string `json:"videoUrl"`
	ContentType int32  `json:"contentType"`
}

type PatternResponse struct {
	Values []PluginPattern `json:"values"`
}

type PluginPattern struct {
	Pattern string         `json:"pattern"`
	Plugin  plugins.Plugin `json:"-"`
}

var registerPlugins = []plugins.Plugin{
	plugins.Facebook{},
	plugins.Instagram{},
	plugins.Tiktok{},
}

func getPatterns() []PluginPattern {
	var patterns []PluginPattern

	for _, plugin := range registerPlugins {
		var patternStrings = plugin.GetPattern()
		for _, patternString := range patternStrings {
			patterns = append(patterns, PluginPattern{
				Pattern: patternString,
				Plugin:  plugin,
			})
		}
	}

	return patterns
}

func getPlugin(patterns []PluginPattern, url string) *plugins.Plugin {
	for _, pattern := range patterns {
		regex, _ := regexp.Compile(pattern.Pattern)
		match := regex.MatchString(url)
		if match {
			return &pattern.Plugin
		}
	}

	return nil
}

func main() {
	e := echo.New()

	var patterns = getPatterns()

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/get-patterns", func(context echo.Context) error {
		return context.JSON(http.StatusOK, PatternResponse{
			Values: patterns,
		})
	})

	e.POST("/get-content", func(context echo.Context) error {
		clientRequest := new(ClientRequest)

		_ = context.Bind(clientRequest)

		var plugin = getPlugin(patterns, clientRequest.Url)
		if plugin != nil {
			var httpClient = http.Client{
				Timeout: time.Second * 10,
			}

			req, _ := http.NewRequest("GET", clientRequest.Url, nil)
			req.Header.Add("cache-control","no-cache")
			req.Header.Add("user-agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

			response, _ := httpClient.Do(req)

			c, _ := ioutil.ReadAll(response.Body)

			content := string(c)

			clientResponse := new(ClientResponse)

			var title = (*plugin).GetTitle(content)
			var thumbnail = (*plugin).GetThumbnail(content)
			var videoData = (*plugin).GetVideoData(content)

			clientResponse.Type = (*plugin).GetType()

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

			if videoData != nil {
				clientResponse.VideoUrl = (*videoData).VideoUrl
				clientResponse.ContentType = (*videoData).ContentType
			} else {
				clientResponse.VideoUrl = ""
				clientResponse.ContentType = -1
			}

			clientResponse.Url = clientRequest.Url

			return context.JSON(http.StatusOK, clientResponse)
		}

		return echo.NewHTTPError(http.StatusNotFound, "Unable to find plugin.")
	})

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Works!!!!!!!")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
