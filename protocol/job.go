package protocol

type (
	Job struct {
		Id                  uint   `json:"id"`
		Name                string `json:"name"`
		LocationID          uint   `json:"location_id"`
		Description         string `json:"description"`
		BasicPaymentPerHour uint   `json:"basic_payment_per_hour"`
	}
	JobServices interface {
		// JobInfo methods
		GetJobs() []*Job
		FindJobByChar(inputChar string) (status bool, res []uint)
		FindJobByID(inputID uint) (error, Job)
		NewJob(inputJob Job) error
		EditJob(ID uint, newJob Job) error
		DeleteJob(IDS []uint) (resDel []uint)
	}
)
