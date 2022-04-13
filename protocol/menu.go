package protocol

type (
	MenuDetail struct {
		Title       string
		Description string
		Route       string
		Type        uint
	}
	MenuGroup struct {
		GroupName string
		Row       []MenuDetail
	}
	Menu struct {
		Group []MenuGroup
	}
)
