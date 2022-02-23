package job

import (
	"strings"
	"tel-note/internal/protocol"
)

type storageMemory protocol.JobStorage

func (AllJob *storageMemory) FindJobByChar(inputChar string) (status bool, res []uint) {
	for _, data := range AllJob.JobData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (AllJob *storageMemory) FindJobByID(inputID uint) (bool, protocol.Job) {
	for _, data := range AllJob.JobData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.Job{}
}

func (AllJob *storageMemory) NewJob(inputJob protocol.Job) bool {
	var LastID uint
	for _, data := range AllJob.JobData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID = +1
	result := protocol.Job{
		Id:                  LastID,
		Name:                inputJob.Name,
		LocationID:          inputJob.LocationID,
		Description:         inputJob.Description,
		BasicPaymentPerHour: inputJob.BasicPaymentPerHour,
	}
	AllJob.JobData = append(AllJob.JobData, &result)
	return true
}

func (AllJob *storageMemory) EditJob(ID uint, newJob protocol.Job) bool {
	for index, data := range AllJob.JobData {
		if data.Id == ID {
			//TODO::: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newJob.Name != "" {
				(AllJob.JobData)[index].Name = newJob.Name
			}
			if newJob.LocationID != 0 {
				(AllJob.JobData)[index].LocationID = newJob.LocationID
			}
			if newJob.Description != "" {
				(AllJob.JobData)[index].Description = newJob.Description
			}
			if newJob.BasicPaymentPerHour != 0 {
				(AllJob.JobData)[index].BasicPaymentPerHour = newJob.BasicPaymentPerHour
			}
			return true
		}
	}
	return false
}

func (AllJob *storageMemory) DeleteJob(IDS []uint) (resDel []uint) {
	for index, id := range IDS {
		for _, data := range AllJob.JobData {
			if data.Id == id {
				AllJob.JobData = append((AllJob.JobData)[:index], (AllJob.JobData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
