package storage

type (
	City struct {
		Id   uint
		Name string
	}
	AllCities struct {
		CityData []*City
	}
)
