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
		NewSex(inputSex string) (bool, SexStorage)
		EditSex(ID uint8, newSex string) bool
		DeleteSex(ID uint8) bool
		FindSexByID(ID uint8) (bool, Sex)
	}
)
