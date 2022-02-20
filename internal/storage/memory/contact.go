package memory

type (
	Contact struct {
		Id uint `json:"id"`
		//TODO::: Person and job must be ID not struct, this way when you modify basic info, contact struct element not edit, because they store data not relation
		*Person
		*JobInfo
		Tel         string `json:"tel,omitempty"`
		Cellphone   string `json:"cellphone,omitempty"`
		Description string `json:"description,omitempty"`
	}
)
