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
		// JobInfo methods
		FindJobByChar(inputChar string) (status bool, res []uint)
		FindJobByID(inputID uint) (bool, Job)
		NewJob(inputJob Job) bool
		EditJob(ID uint, newJob Job) bool
		DeleteJob(IDS []uint) (resDel []uint)
	}
)
