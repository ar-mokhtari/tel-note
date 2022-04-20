//go:build !mysql || !mongodb

package sex

import (
	"errors"
	"tel-note/protocol"
)

type (
	storageMemory struct {
		SexData []*protocol.Sex
	}
)

var (
	storage storageMemory
	_       protocol.SexServices = &storage
)

func (allSex *storageMemory) GetSex() []*protocol.Sex {
	return allSex.SexData
}

func (allSex *storageMemory) NewSex(inputSex protocol.Sex) error {
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
	return nil
}

func (allSex *storageMemory) EditSex(newSex protocol.Sex) error {
	for index, data := range allSex.SexData {
		if data.Id == newSex.Id {
			if newSex.Name != "" {
				(allSex.SexData)[index].Name = newSex.Name
			}
			return nil
		}
	}
	return errors.New("not found")
}

func (allSex *storageMemory) DeleteSex(ID uint8) error {
	for index, data := range allSex.SexData {
		if data.Id == ID {
			allSex.SexData = append((allSex.SexData)[:index], (allSex.SexData)[index+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (allSex *storageMemory) FindSexByID(ID uint8) protocol.Sex {
	for _, data := range allSex.SexData {
		if data.Id == ID {
			return *data
		}
	}
	return protocol.Sex{}
}
