package config

type (
	Greeting struct {
		General     string `json:"general,omitempty"`
		Description string `json:"description,omitempty"`
	}
	Role           int
	ResponseStatus struct {
		State  bool
		String string
	}
)

const (
	AdminState Role = iota + 1
	Manager
	Basic
	AdminString = "admin"
	OkStatus    = "yes"
)
