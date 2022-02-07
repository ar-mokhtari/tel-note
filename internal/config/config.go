package config

type (
	Greeting struct {
		General     string `json:"general,omitempty"`
		Description string `json:"description,omitempty"`
	}
	Role int
)

const (
	Admin Role = iota + 1
	manager
	basic
	OkStatus = "yes"
)
