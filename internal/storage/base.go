package storage

type (
	AllDataTool interface {
		GetANewCodeID()
	}
	AllData struct {
		ContactData []*Contact
		CityData    []*City
		jobData     []*jobInfo
	}
)
