package plugins

type VideoData struct {
	VideoUrl    string
	ContentType int32
}

type Plugin interface {
	GetPattern() []string
	GetType() int32
	GetTitle(content string) *string
	GetThumbnail(content string) *string
	GetVideoData(content string) *VideoData
}
