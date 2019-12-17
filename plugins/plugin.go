package plugins

type Plugin interface {
	GetPattern() []string
	GetType() int32
	GetTitle(content string) *string
	GetThumbnail(content string) *string
	GetVideoUrl(content string) *string
}
