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

func (r Facebook) GetType() int32 {
	return 0
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

func (r Facebook) GetVideoData(content string) *VideoData {
	regex1, _ := regexp.Compile(`hd_src:"([^"]+)"`)
	match1 := regex1.FindStringSubmatch(content)
	if match1 != nil && len(match1) > 1 {
		return &VideoData{
			VideoUrl:    match1[1],
			ContentType: 0,
		}
	}

	regex2, _ := regexp.Compile(`property="og:video" content="([^"]+)"`)
	match2 := regex2.FindStringSubmatch(content)
	if match2 != nil && len(match2) > 1 {
		var str = match2[1]
		str = strings.Replace(str, "&amp;", "&", -1)
		return &VideoData{
			VideoUrl:    str,
			ContentType: 0,
		}
	}

	return nil
}
