package plugins

import (
	"regexp"
	"strings"
)

type Facebook struct {
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
		var txt = match[1]
		txt = strings.Replace(txt, "&amp;", "&", -1)
		return &txt
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
