package plugins

type Plugin interface {
	GetPattern() []string
	GetTitle(content string) *string
	GetThumbnail(content string) *string
	GetVideoUrl(content string) *string
}
