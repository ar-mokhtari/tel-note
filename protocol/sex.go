package protocol

type (
	Sex struct {
		Id   byte   `json:"id"`
		Name string `json:"name"`
	}
	SexServices interface {
		GetSex() []*Sex
		NewSex(inputSex Sex) error
		EditSex(newSex Sex) error
		DeleteSex(ID byte) error
		FindSexByID(ID byte) Sex
	}
)
