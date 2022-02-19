package memory

type (
	JobInfo struct {
		Id                  uint
		Name                string
		Location            *City
		Description         string
		BasicPaymentPerHour uint
	}
)
