package person

import (
	"strings"
	"tel-note/protocol"
	"time"
)

type storageMemory protocol.PersonStorage

func (AllPerson *storageMemory) GetPersons() protocol.PersonStorage {
	return protocol.PersonStorage(*AllPerson)
}

func (AllPerson *storageMemory) FindPersonByChar(inputChar string) (status bool, res protocol.PersonStorage) {
	for _, data := range AllPerson.PersonData {
		if strings.Contains(
			strings.ToLower(data.FirstName+data.LastName),
			strings.ToLower(inputChar),
		) {
			res.PersonData = append(res.PersonData, data)
			status = true
		}
	}
	return status, res
}

func (AllPerson *storageMemory) FindPersonByID(inputID uint) (bool, protocol.Person) {
	for _, data := range AllPerson.PersonData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.Person{}
}

func (AllPerson *storageMemory) NewPerson(inputPerson protocol.Person) bool {
	var LastID uint
	for _, data := range AllPerson.PersonData {
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
	AllPerson.PersonData = append(AllPerson.PersonData, &result)
	return true
}

func (AllPerson *storageMemory) EditPerson(ID uint, newPerson protocol.Person) bool {
	for index, data := range AllPerson.PersonData {
		if data.Id == ID {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newPerson.FirstName != "" {
				(AllPerson.PersonData)[index].FirstName = newPerson.FirstName
			}
			if newPerson.LastName != "" {
				(AllPerson.PersonData)[index].LastName = newPerson.LastName
			}
			if newPerson.Description != "" {
				(AllPerson.PersonData)[index].Description = newPerson.Description
			}
			if !(newPerson.DOB).IsZero() {
				(AllPerson.PersonData)[index].DOB = newPerson.DOB
			}
			if newPerson.BirthLocationID != 0 {
				(AllPerson.PersonData)[index].BirthLocationID = newPerson.BirthLocationID
			}
			if newPerson.NationalCode != "" {
				(AllPerson.PersonData)[index].NationalCode = newPerson.NationalCode
			}
			if newPerson.GenderID != 0 {
				(AllPerson.PersonData)[index].GenderID = newPerson.GenderID
			}
			if newPerson.FirstName != "" ||
				newPerson.LastName != "" ||
				newPerson.Description != "" ||
				!(newPerson.DOB).IsZero() ||
				newPerson.BirthLocationID != 0 ||
				newPerson.NationalCode != "" ||
				newPerson.GenderID != 0 {
				(AllPerson.PersonData)[index].UpdateAt = time.Now()
			}
			return true
		}
	}
	return false
}

func (AllPerson *storageMemory) DeletePerson(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range AllPerson.PersonData {
			if data.Id == id {
				AllPerson.PersonData = append((AllPerson.PersonData)[:index], (AllPerson.PersonData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
