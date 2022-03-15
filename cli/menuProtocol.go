package cli

type (
	MenuDetail struct {
		Title       string
		Description string
	}
	MenuGroup struct {
		GroupName string
		Row       []MenuDetail
	}
	Menu struct {
		Group []MenuGroup
	}
)
