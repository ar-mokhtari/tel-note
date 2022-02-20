package memory

import (
	"strings"
)

type (
	AllDataTool interface {
		// Contact methods
		AddContact(inputContact Contact) (bool, AllData)
		FindContactByID(id uint) (Contact, bool)
		FindContactByChar(insertChar string) (AllData, uint)
		EditContactByID(newData Contact, ID uint) bool
		DeleteContactByID(ID uint) bool
		DeleteAll() bool
		// City methods
		NewCity(CityName string) bool
		EditCityByID(ID uint, NewCityName string) bool
		DeleteCityByID(ID uint) bool
		// JobInfo methods
		NewJob(jobName string) bool
		EditJobByID(ID uint, NewJobName string) bool
		DeleteJobByID(ID uint) bool
		FindJobById(ID uint) string
		//other methods
		//TODO::: don't need this method for interface, it's internal in memory/base.go:
		//FindGenderNameById(ID uint8) string
		NewSex(sexName string) bool
		EditSex(ID uint8, newName string) bool
		DeleteSexByID(ID uint8) bool
	}

	AllData struct {
		ContactData []*Contact
		JobData     []*JobInfo
		PersonData  []*Person
		CityData    []*City
		SexData     []*Sex
	}
)

// Contact methods

func (MainData *AllData) AddContact(inputContact Contact) (bool, AllData) {
	var (
		lastID uint
	)
	for _, data := range MainData.ContactData {
		if data.Id > lastID {
			lastID = data.Id
		}
	}
	lastID += 1
	//find job name by ID to return for end user
	inputContact.JobInfo.Name = MainData.FindJobById(inputContact.JobInfo.Id)
	//find gender name by ID to return for end user
	inputContact.Gender.Name = MainData.FindGenderNameById(inputContact.Gender.Id)
	//marge inputs to create a contact
	result := &Contact{
		Id: lastID,
		Person: &Person{
			FirstName: inputContact.FirstName,
			LastName:  inputContact.LastName,
			Gender: &Sex{
				Id:   inputContact.Gender.Id,
				Name: inputContact.Gender.Name,
			},
		},
		Tel:         inputContact.Tel,
		Cellphone:   inputContact.Cellphone,
		Description: inputContact.Description,
		JobInfo: &JobInfo{
			Id:   inputContact.JobInfo.Id,
			Name: inputContact.JobInfo.Name,
		},
	}
	MainData.ContactData = append(MainData.ContactData, result)
	return true, *MainData
}

func (MainData *AllData) FindContactByID(id uint) (Contact, bool) {
	var (
		data Contact
	)
	for _, data := range MainData.ContactData {
		if data.Id == id {
			return *data, true
			break
		}
	}
	return data, false
}

func (MainData *AllData) FindContactByChar(insertChar string) (AllData, uint) {
	var (
		counter    uint
		resultData AllData
	)
	for _, data := range MainData.ContactData {
		if strings.Contains(data.FirstName, insertChar) {
			counter += 1
			resultData.ContactData = append(resultData.ContactData, data)
		}
	}
	return resultData, counter
}

func (MainData *AllData) EditContactByID(newData Contact, ID uint) bool {
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newData.FirstName != "" {
				(MainData.ContactData)[index].FirstName = newData.FirstName
			}
			if newData.LastName != "" {
				(MainData.ContactData)[index].LastName = newData.LastName
			}
			if newData.Description != "" {
				(MainData.ContactData)[index].Description = newData.Description
			}
			if newData.Tel != "" {
				(MainData.ContactData)[index].Tel = newData.Tel
			}
			if newData.Cellphone != "" {
				(MainData.ContactData)[index].Cellphone = newData.Cellphone
			}
			if newData.Gender.Id != 0 {
				(MainData.ContactData)[index].Gender.Id = newData.Gender.Id
				(MainData.ContactData)[index].Gender.Name = MainData.FindGenderNameById(newData.Gender.Id)
			}
			if newData.JobInfo.Id != 0 {
				(MainData.ContactData)[index].JobInfo.Id = newData.JobInfo.Id
				(MainData.ContactData)[index].JobInfo.Name = MainData.FindJobById(newData.JobInfo.Id)
			}
			return true
		}
	}
	return false
}

func (MainData *AllData) DeleteContactByID(ID uint) bool {
	for index, data := range MainData.ContactData {
		if data.Id == ID {
			MainData.ContactData = append(MainData.ContactData[:index], MainData.ContactData[index+1:]...)
			return true
		}
	}
	return false
}

func (MainData *AllData) DeleteAll() bool {
	MainData.ContactData = MainData.ContactData[0:0]
	return true
}

// City methods

func (MainData *AllData) NewCity(CityName string) bool {
	var LastID uint
	for _, data := range MainData.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := City{
		Id:   uint(LastID) + 1,
		Name: CityName,
	}
	MainData.CityData = append(MainData.CityData, &result)
	return true
}

func (MainData *AllData) EditCityByID(ID uint, NewCityName string) bool {
	for _, data := range MainData.CityData {
		if data.Id == ID {
			data.Name = NewCityName
			return true
		}
	}
	return false
}

func (MainData *AllData) DeleteCityByID(ID uint) bool {
	for index, data := range MainData.CityData {
		if data.Id == ID {
			MainData.CityData = append(MainData.CityData[:index], MainData.CityData[index+1:]...)
			return true
		}
	}
	return false
}

// Job methods

func (MainData *AllData) NewJob(jobName string) bool {
	var LastID uint
	for _, data := range MainData.JobData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := JobInfo{
		Id:   uint(LastID) + 1,
		Name: jobName,
	}
	MainData.JobData = append(MainData.JobData, &result)
	return true
}

func (MainData *AllData) EditJobByID(ID uint, NewJobName string) bool {
	for _, data := range MainData.JobData {
		if data.Id == ID {
			data.Name = NewJobName
			return true
		}
	}
	return false
}

func (MainData *AllData) DeleteJobByID(ID uint) bool {
	for index, data := range MainData.JobData {
		if data.Id == ID {
			MainData.JobData = append(MainData.JobData[:index], MainData.JobData[index+1:]...)
			return true
		}
	}
	return false
}

func (MainData *AllData) FindJobById(ID uint) string {
	for _, data := range MainData.JobData {
		if ID == data.Id {
			return data.Name
		}
	}
	return ""
}

// other methods

func (MainData *AllData) FindGenderNameById(ID uint8) string {
	for _, data := range MainData.SexData {
		if ID == data.Id {
			return data.Name
		}
	}
	return ""
}

func (MainData *AllData) NewSex(sexName string) bool {
	var LastID uint8
	for _, data := range MainData.SexData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	result := Sex{
		Id:   uint8(LastID) + 1,
		Name: sexName,
	}
	MainData.SexData = append(MainData.SexData, &result)
	return true
}

func (MainData *AllData) EditSex(ID uint8, newName string) bool {
	for _, data := range MainData.SexData {
		if data.Id == ID {
			data.Name = newName
			return true
		}
	}
	return true
}

func (MainData *AllData) DeleteSexByID(ID uint8) bool {
	for index, data := range MainData.SexData {
		if data.Id == ID {
			MainData.SexData = append(MainData.SexData[:index], MainData.SexData[index+1:]...)
			return true
		}
	}
	return false
}
