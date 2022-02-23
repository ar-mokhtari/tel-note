package sex

import (
	"tel-note/internal/protocol"
)

type storageMemory protocol.SexStorage

func (AllSex *storageMemory) NewSex(inputSex string) bool {
	var LastID uint8
	for _, data := range AllSex.SexData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID = +1
	result := protocol.Sex{
		Id:   LastID,
		Name: inputSex,
	}
	AllSex.SexData = append(AllSex.SexData, &result)
	return true
}

func (AllSex *storageMemory) EditSex(ID uint8, newSex string) bool {
	for index, data := range AllSex.SexData {
		if data.Id == ID {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newSex != "" {
				(AllSex.SexData)[index].Name = newSex
			}
			return true
		}
	}
	return false
}

func (AllSex *storageMemory) DeleteSex(ID uint8) bool {
	for index, data := range AllSex.SexData {
		if data.Id == ID {
			AllSex.SexData = append((AllSex.SexData)[:index], (AllSex.SexData)[index+1:]...)
			return true
		}
	}
	return false
}
