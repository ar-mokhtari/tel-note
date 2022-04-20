package protocol

type (
	Sex struct {
		Id   uint8  `json:"id"`
		Name string `json:"name"`
	}
	SexServices interface {
		GetSex() []*Sex
		NewSex(inputSex Sex) error
		EditSex(newSex Sex) error
		DeleteSex(ID uint8) error
		FindSexByID(ID uint8) Sex
	}
)
