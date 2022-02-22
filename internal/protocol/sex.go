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
		NewSex(inputSex string) bool
		EditSex(ID uint8, newSex string) bool
		DeleteSex(ID uint8) bool
	}
)
