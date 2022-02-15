package memory

type (
	Contact struct {
		Id uint `json:"id"`
		*Person
		*jobInfo
		Tel         string `json:"tel,omitempty"`
		Cellphone   string `json:"cellphone,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
