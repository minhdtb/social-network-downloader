package plugins

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
	return nil
}

func (r Instagram) GetThumbnail(content string) *string {
	return nil
}

func (r Instagram) GetVideoUrl(content string) *string {
	return nil
}
