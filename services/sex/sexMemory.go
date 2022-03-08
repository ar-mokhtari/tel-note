package sex

import (
	"tel-note/protocol"
)

type storageMemory protocol.SexStorage

func (AllSex *storageMemory) GetSex() protocol.SexStorage {
	return protocol.SexStorage(*AllSex)
}

func (AllSex *storageMemory) NewSex(inputSex protocol.Sex) bool {
	var LastID uint8
	for _, data := range AllSex.SexData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID += 1
	result := protocol.Sex{
		Id:   LastID,
		Name: inputSex.Name,
	}
	AllSex.SexData = append(AllSex.SexData, &result)
	return true
}

//TODO::: ForEditStrategy,AllEntity(ies)CanInputProtocolArgumentWithIdentityFieldAndDon'tNeedThisFieldSeparately
func (AllSex *storageMemory) EditSex(newSex protocol.Sex) bool {
	for index, data := range AllSex.SexData {
		if data.Id == newSex.Id {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newSex.Name != "" {
				(AllSex.SexData)[index].Name = newSex.Name
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

func (AllSex *storageMemory) FindSexByID(ID uint8) (bool, protocol.Sex) {
	for _, data := range AllSex.SexData {
		if data.Id == ID {
			return true, *data
		}
	}
	return false, protocol.Sex{}
}
