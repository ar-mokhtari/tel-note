package job

import "tel-note/protocol"

func GetJobs() []*protocol.Job {
	return storage.GetJobs()
}
