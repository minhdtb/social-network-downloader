package plugins

import (
	"regexp"
	"strings"
)

type Facebook struct {
}

func (r Facebook) GetPattern() []string {
	return []string{
		`(?:http:\/\/)?(?:www\.)?facebook\.com\/([a-zA-Z0-9_.-]+)\/posts\/`,
		`(?:http:\/\/)?(?:www\.)?facebook\.com\/([a-zA-Z0-9_.-]+)\/videos\/`,
	}
}

func (r Facebook) GetTitle(content string) *string {
	regex, _ := regexp.Compile(`title id="pageTitle">(.+?)</title>`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		return &match[1]
	}

	return nil
}

func (r Facebook) GetThumbnail(content string) *string {
	regex, _ := regexp.Compile(`property="og:image" content="([^"]+)"`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		var str = match[1]
		str = strings.Replace(str, "&amp;", "&", -1)
		return &str
	}

	return nil
}

func (r Facebook) GetVideoUrl(content string) *string {
	regex, _ := regexp.Compile(`hd_src:"([^"]+)"`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		return &match[1]
	}

	return nil
}
