package jobInfo

import (
	"tel-note/internal/config"
	"tel-note/internal/storage/memory"
)

func NewJob(jobName string) config.ResponseStatus {
	if memory.AllDataTool.NewJob(&config.MainData, jobName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}

}

func EditJobInfoByID(ID uint, NewJobInfoName string) config.ResponseStatus {
	if memory.AllDataTool.EditJobByID(&config.MainData, ID, NewJobInfoName) {
		return config.ResponseStatus{State: true}
	}
	return config.ResponseStatus{State: false}
}

func DeleteJobByID(ID uint) *config.ResponseStatus {
	if memory.AllDataTool.DeleteJobByID(&config.MainData, ID) {
		return &config.ResponseStatus{State: true}
	}
	return &config.ResponseStatus{State: false}
}
