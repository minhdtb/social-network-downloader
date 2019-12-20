package plugins

import (
	"regexp"
	"strings"
)

type Instagram struct {
}

func (r Instagram) GetPattern() []string {
	return []string{
		`(?:http:\/\/)?(?:www\.)?instagram\.com\/p\/`,
	}
}

func (r Instagram) GetType() int32 {
	return 1
}

func (r Instagram) GetTitle(content string) *string {
	regex2, _ := regexp.Compile(`property="og:title" content="([^"]+)"`)
	match := regex2.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		return &match[1]
	}

	return nil
}

func (r Instagram) GetThumbnail(content string) *string {
	regex2, _ := regexp.Compile(`property="og:image" content="([^"]+)"`)
	match := regex2.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		return &match[1]
	}

	return nil
}

func (r Instagram) GetVideoUrl(content string) *string {
	regex2, _ := regexp.Compile(`property="og:video" content="([^"]+)"`)
	match := regex2.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		var str = match[1]
		str = strings.Replace(str, "&amp;", "&", -1)
		return &str
	}

	return nil
}
