package job

import "tel-note/protocol"

func GetJobs() protocol.JobStorage {
	return storage.GetJobs()
}
