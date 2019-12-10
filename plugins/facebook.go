package plugins

import "regexp"

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
