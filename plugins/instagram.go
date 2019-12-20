package plugins

import (
	"encoding/json"
	"regexp"
)

type ShortcodeMedia struct {
	Id         string `json:"id"`
	DisplayUrl string `json:"display_url"`
	VideoUrl   string `json:"video_url"`
	IsVideo    bool   `json:"is_video"`
}

type GraphQl struct {
	ShortcodeMedia ShortcodeMedia `json:"shortcode_media"`
}

type PostPageItem struct {
	GraphQl GraphQl `json:"graphql"`
}

type EntryData struct {
	PostPage []PostPageItem
}

type InstagramData struct {
	EntryData EntryData `json:"entry_data"`
}

type Instagram struct {
}

func getData(content string) *ShortcodeMedia {
	regex, _ := regexp.Compile(`window._sharedData = ((?s).*)};</script>`)
	match := regex.FindStringSubmatch(content)
	if match != nil && len(match) > 1 {
		var data InstagramData
		var str = match[1] + "}"
		err := json.Unmarshal([]byte(str), &data)
		if err != nil {
			return nil
		}

		if len(data.EntryData.PostPage) > 0 {
			return &(data.EntryData.PostPage[0].GraphQl.ShortcodeMedia)
		}
	}

	return nil
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

func (r Instagram) GetVideoUrl(content string) *VideoData {
	data := getData(content)
	if data != nil {
		if data.IsVideo {
			return &VideoData{
				VideoUrl:    data.VideoUrl,
				ContentType: 0,
			}
		} else {
			return &VideoData{
				VideoUrl:    data.DisplayUrl,
				ContentType: 1,
			}
		}
	}

	return nil
}
