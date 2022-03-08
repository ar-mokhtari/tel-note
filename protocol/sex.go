package protocol

type (
	Sex struct {
		Id   uint8
		Name string
	}
	SexStorage struct {
		SexData []*Sex
	}
	SexServices interface {
		GetSex() SexStorage
		NewSex(inputSex Sex) bool
		EditSex(newSex Sex) bool
		DeleteSex(ID uint8) bool
		FindSexByID(ID uint8) (bool, Sex)
	}
)
