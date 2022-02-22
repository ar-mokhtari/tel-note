package protocol

type (
	Job struct {
		Id                  uint
		Name                string
		LocationID          uint
		Description         string
		BasicPaymentPerHour uint
	}
	JobStorage struct {
		JobData []*Job
	}
	JobServices interface {
		NewJob(job Job) bool
		FindJobByID(inputID uint) (bool, Job)
	}
)
