package plugins

import (
	"regexp"
	"strings"
)

type Tiktok struct {
}

func (r Tiktok) GetPattern() []string {
	return []string{
		`(?:http:\/\/)?(?:www\.)?tiktok\.com\/(.+)\/video\/`,
	}
}

func (r Tiktok) GetType() int32 {
	return 2
}

func (r Tiktok) GetTitle(content string) *string {
	regex, _ := regexp.Compile(`property="og:title" content="([^"]+)"`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		return &match[1]
	}

	return nil
}

func (r Tiktok) GetThumbnail(content string) *string {
	regex, _ := regexp.Compile(`property="og:image" content="([^"]+)"`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		var str = match[1]
		str = strings.Replace(str, "&amp;", "&", -1)
		return &str
	}

	return nil
}

func (r Tiktok) GetVideoData(content string) *VideoData {
	regex, _ := regexp.Compile(`<video playsinline="" loop="" pageType="0" src="([^"]+)"`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		return &VideoData{
			VideoUrl:    match[1],
			ContentType: 0,
		}
	}

	return nil
}
