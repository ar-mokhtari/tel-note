package protocol

type (
	MenuDetail struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Route       string `json:"route"`
		Type        uint   `json:"type"`
	}
	MenuGroup struct {
		GroupName string
		Row       []MenuDetail
	}
	Menu struct {
		Group []MenuGroup
	}
)
