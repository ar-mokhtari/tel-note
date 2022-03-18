package sex

import (
	"tel-note/protocol"
)

type (
	storageMemory struct {
		SexData []*protocol.Sex
	}
)

func (allSex *storageMemory) GetSex() []*protocol.Sex {
	return allSex.SexData
}

func (allSex *storageMemory) NewSex(inputSex protocol.Sex) bool {
	var LastID uint8
	for _, data := range allSex.SexData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID += 1
	result := protocol.Sex{
		Id:   LastID,
		Name: inputSex.Name,
	}
	allSex.SexData = append(allSex.SexData, &result)
	return true
}

//TODO::: ForEditStrategy,AllEntity(ies)CanInputProtocolArgumentWithIdentityFieldAndDon'tNeedThisFieldSeparately
func (allSex *storageMemory) EditSex(newSex protocol.Sex) bool {
	for index, data := range allSex.SexData {
		if data.Id == newSex.Id {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newSex.Name != "" {
				(allSex.SexData)[index].Name = newSex.Name
			}
			return true
		}
	}
	return false
}

func (allSex *storageMemory) DeleteSex(ID uint8) bool {
	for index, data := range allSex.SexData {
		if data.Id == ID {
			allSex.SexData = append((allSex.SexData)[:index], (allSex.SexData)[index+1:]...)
			return true
		}
	}
	return false
}

func (allSex *storageMemory) FindSexByID(ID uint8) (bool, protocol.Sex) {
	for _, data := range allSex.SexData {
		if data.Id == ID {
			return true, *data
		}
	}
	return false, protocol.Sex{}
}
