package person

import (
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

func (allPerson *storageMemory) GetPersons() []*protocol.Person {
	return allPerson.PersonData
}

func (allPerson *storageMemory) FindPersonByChar(inputChar string) (status bool, res []*protocol.Person) {
	for _, data := range allPerson.PersonData {
		if strings.Contains(
			strings.ToLower(data.FirstName+data.LastName),
			strings.ToLower(inputChar),
		) {
			res = append(res, data)
			status = true
		}
	}
	return status, res
}

func (allPerson *storageMemory) FindPersonByID(inputID uint) (bool, protocol.Person) {
	for _, data := range allPerson.PersonData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.Person{}
}

func (allPerson *storageMemory) NewPerson(inputPerson protocol.Person) bool {
	var LastID uint
	for _, data := range allPerson.PersonData {
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
	allPerson.PersonData = append(allPerson.PersonData, &result)
	return true
}

func (allPerson *storageMemory) EditPerson(ID uint, newPerson protocol.Person) bool {
	for index, data := range allPerson.PersonData {
		if data.Id == ID {
			if newPerson.FirstName != "" {
				(allPerson.PersonData)[index].FirstName = newPerson.FirstName
			}
			if newPerson.LastName != "" {
				(allPerson.PersonData)[index].LastName = newPerson.LastName
			}
			if newPerson.Description != "" {
				(allPerson.PersonData)[index].Description = newPerson.Description
			}
			if !(newPerson.DOB).IsZero() {
				(allPerson.PersonData)[index].DOB = newPerson.DOB
			}
			if newPerson.BirthLocationID != 0 {
				(allPerson.PersonData)[index].BirthLocationID = newPerson.BirthLocationID
			}
			if newPerson.NationalCode != "" {
				(allPerson.PersonData)[index].NationalCode = newPerson.NationalCode
			}
			if newPerson.GenderID != 0 {
				(allPerson.PersonData)[index].GenderID = newPerson.GenderID
			}
			if newPerson.FirstName != "" ||
				newPerson.LastName != "" ||
				newPerson.Description != "" ||
				!(newPerson.DOB).IsZero() ||
				newPerson.BirthLocationID != 0 ||
				newPerson.NationalCode != "" ||
				newPerson.GenderID != 0 {
				(allPerson.PersonData)[index].UpdateAt = time.Now()
			}
			return true
		}
	}
	return false
}

func (allPerson *storageMemory) DeletePerson(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range allPerson.PersonData {
			if data.Id == id {
				allPerson.PersonData = append((allPerson.PersonData)[:index], (allPerson.PersonData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
