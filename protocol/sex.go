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
		NewSex(inputSex Sex) (bool, SexStorage)
		EditSex(newSex Sex) bool
		DeleteSex(ID uint8) bool
		FindSexByID(ID uint8) (bool, Sex)
	}
)
