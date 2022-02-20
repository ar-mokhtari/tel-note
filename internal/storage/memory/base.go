package memory

type (
	AllDataTool interface {
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
		JobData    []*JobInfo
		PersonData []*Person
		CityData   []*City
		SexData    []*Sex
	}
)

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
