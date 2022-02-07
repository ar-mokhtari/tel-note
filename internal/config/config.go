package config

type (
	Greeting struct {
		General     string `json:"general,omitempty"`
		Description string `json:"description,omitempty"`
	}
	Role int
)

const (
	AdminState  Role = iota + 1
	AdminString      = "admin"
	manager
	basic
	OkStatus = "yes"
)
