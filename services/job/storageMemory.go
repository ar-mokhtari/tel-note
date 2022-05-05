//go:build !mysql || !mongodb

package job

import (
	"errors"
	"github.com/ar-mokhtari/tel-note/protocol"
	"strings"
)

type storageMemory struct {
	JobData []*protocol.Job
}

var (
	storage storageMemory
	_       protocol.JobServices = &storage
)

func (sm *storageMemory) GetJobs() []*protocol.Job {
	return sm.JobData
}

func (sm *storageMemory) FindJobByChar(inputChar string) (status bool, res []uint) {
	for _, data := range sm.JobData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (sm *storageMemory) FindJobByID(inputID uint) (error, protocol.Job) {
	for _, data := range sm.JobData {
		if data.Id == inputID {
			return nil, *data
		}
	}
	return errors.New("not found"), protocol.Job{}
}

func (sm *storageMemory) NewJob(inputJob protocol.Job) error {
	var LastID uint
	for _, data := range sm.JobData {
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
	sm.JobData = append(sm.JobData, &result)
	return nil
}

func (sm *storageMemory) EditJob(ID uint, newJob protocol.Job) error {
	for index, data := range sm.JobData {
		if data.Id == ID {
			if newJob.Name != "" {
				(sm.JobData)[index].Name = newJob.Name
			}
			if newJob.LocationID != 0 {
				(sm.JobData)[index].LocationID = newJob.LocationID
			}
			if newJob.Description != "" {
				(sm.JobData)[index].Description = newJob.Description
			}
			if newJob.BasicPaymentPerHour != 0 {
				(sm.JobData)[index].BasicPaymentPerHour = newJob.BasicPaymentPerHour
			}
			return nil
		}
	}
	return errors.New("job not found")
}

func (sm *storageMemory) DeleteJob(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range sm.JobData {
			if data.Id == id {
				sm.JobData = append((sm.JobData)[:index], (sm.JobData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}
