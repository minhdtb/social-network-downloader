package plugins

type Plugin interface {
	GetTitle(content string) *string
	GetThumbnail(content string) *string
	GetVideoUrl(content string) *string
}
