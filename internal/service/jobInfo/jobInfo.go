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
