package protocol

type (
	Sex struct {
		Id   uint8
		Name string
	}
	SexServices interface {
		GetSex() []*Sex
		NewSex(inputSex Sex) bool
		EditSex(newSex Sex) bool
		DeleteSex(ID uint8) bool
		FindSexByID(ID uint8) (bool, Sex)
	}
)
