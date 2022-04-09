package protocol

type (
	Greeting struct {
		General     string `json:"general,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
