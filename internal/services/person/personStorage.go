package person

import (
	"strings"
	"tel-note/internal/protocol"
)

type Storage struct {
	protocol.PersonStorage
}

func (AllPerson *Storage) FindPersonByChar(inputChar string) (status bool, res []uint) {
	for _, data := range *AllPerson.PersonData {
		if strings.Contains(data.FirstName, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}
