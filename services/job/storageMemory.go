package job

import (
	"strings"
	"tel-note/protocol"
)

type storageMemory struct {
	JobData []*protocol.Job
}

var (
	storage storageMemory
	_       protocol.JobServices = &storage
)

func (allJob *storageMemory) GetJobs() []*protocol.Job {
	return allJob.JobData
}

func (allJob *storageMemory) FindJobByChar(inputChar string) (status bool, res []uint) {
	for _, data := range allJob.JobData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (allJob *storageMemory) FindJobByID(inputID uint) (bool, protocol.Job) {
	for _, data := range allJob.JobData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.Job{}
}

func (allJob *storageMemory) NewJob(inputJob protocol.Job) bool {
	var LastID uint
	for _, data := range allJob.JobData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID += 1
	result := protocol.Job{
		Id:                  LastID,
		Name:                inputJob.Name,
		LocationID:          inputJob.LocationID,
		Description:         inputJob.Description,
		BasicPaymentPerHour: inputJob.BasicPaymentPerHour,
	}
	allJob.JobData = append(allJob.JobData, &result)
	return true
}

func (allJob *storageMemory) EditJob(ID uint, newJob protocol.Job) bool {
	for index, data := range allJob.JobData {
		if data.Id == ID {
			if newJob.Name != "" {
				(allJob.JobData)[index].Name = newJob.Name
			}
			if newJob.LocationID != 0 {
				(allJob.JobData)[index].LocationID = newJob.LocationID
			}
			if newJob.Description != "" {
				(allJob.JobData)[index].Description = newJob.Description
			}
			if newJob.BasicPaymentPerHour != 0 {
				(allJob.JobData)[index].BasicPaymentPerHour = newJob.BasicPaymentPerHour
			}
			return true
		}
	}
	return false
}

func (allJob *storageMemory) DeleteJob(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range allJob.JobData {
			if data.Id == id {
				allJob.JobData = append((allJob.JobData)[:index], (allJob.JobData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
