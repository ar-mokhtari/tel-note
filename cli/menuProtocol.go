package cli

type (
	MenuDetail struct {
		Title       string
		Description string
		Route       string
	}
	MenuGroup struct {
		GroupName string
		Row       []MenuDetail
	}
	Menu struct {
		Group []MenuGroup
	}
)
