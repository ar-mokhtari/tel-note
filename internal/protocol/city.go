package protocol

type (
	City struct {
		Id       uint
		Name     string
		AriaCode string
	}
	CityStorage struct {
		CityData []*City
	}
	CityServices interface {
		NewCity()
	}
)
