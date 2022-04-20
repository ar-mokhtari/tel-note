//go:build !mysql || !mongodb

package person

import (
	"errors"
	"strings"
	"tel-note/protocol"
	"time"
)

type storageMemory struct {
	PersonData []*protocol.Person
}

var (
	storage storageMemory
	_       protocol.PersonServices = &storage
)

func (sm *storageMemory) GetPersons() []*protocol.Person {
	return sm.PersonData
}

func (sm *storageMemory) FindPersonByChar(inputChar string) (res []*protocol.Person) {
	for _, data := range sm.PersonData {
		if strings.Contains(
			strings.ToLower(data.FirstName+data.LastName),
			strings.ToLower(inputChar),
		) {
			res = append(res, data)
		}
	}
	return res
}

func (sm *storageMemory) FindPersonByID(inputID uint) protocol.Person {
	for _, data := range sm.PersonData {
		if data.Id == inputID {
			return *data
		}
	}
	return protocol.Person{}
}

func (sm *storageMemory) NewPerson(inputPerson protocol.Person) error {
	var LastID uint
	for _, data := range sm.PersonData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID += 1
	result := protocol.Person{
		Id:              LastID,
		FirstName:       inputPerson.FirstName,
		LastName:        inputPerson.LastName,
		DOB:             inputPerson.DOB,
		BirthLocationID: inputPerson.BirthLocationID,
		GenderID:        inputPerson.GenderID,
		NationalCode:    inputPerson.NationalCode,
		Description:     inputPerson.Description,
		CreateAt:        time.Now(),
	}
	sm.PersonData = append(sm.PersonData, &result)
	return nil
}

func (sm *storageMemory) EditPerson(ID uint, newPerson protocol.Person) error {
	for index, data := range sm.PersonData {
		if data.Id == ID {
			if newPerson.FirstName != "" {
				(sm.PersonData)[index].FirstName = newPerson.FirstName
			}
			if newPerson.LastName != "" {
				(sm.PersonData)[index].LastName = newPerson.LastName
			}
			if newPerson.Description != "" {
				(sm.PersonData)[index].Description = newPerson.Description
			}
			if !(newPerson.DOB).IsZero() {
				(sm.PersonData)[index].DOB = newPerson.DOB
			}
			if newPerson.BirthLocationID != 0 {
				(sm.PersonData)[index].BirthLocationID = newPerson.BirthLocationID
			}
			if newPerson.NationalCode != "" {
				(sm.PersonData)[index].NationalCode = newPerson.NationalCode
			}
			if newPerson.GenderID != 0 {
				(sm.PersonData)[index].GenderID = newPerson.GenderID
			}
			if newPerson.FirstName != "" ||
				newPerson.LastName != "" ||
				newPerson.Description != "" ||
				!(newPerson.DOB).IsZero() ||
				newPerson.BirthLocationID != 0 ||
				newPerson.NationalCode != "" ||
				newPerson.GenderID != 0 {
				(sm.PersonData)[index].UpdateAt = time.Now()
			}
			return nil
		}
	}
	return errors.New("not found")
}

func (sm *storageMemory) DeletePerson(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range sm.PersonData {
			if data.Id == id {
				sm.PersonData = append((sm.PersonData)[:index], (sm.PersonData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
