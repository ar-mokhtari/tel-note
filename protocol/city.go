package protocol

type (
	City struct {
		Id       uint
		Name     string
		AriaCode string
		Lat      float64
		Lng      float64
	}
	CityStorage struct {
		CityData []*City
	}
	CityServices interface {
		//city methods
		GetCities() CityStorage
		FindCityByChar(inputChar string) (status bool, res []uint)
		FindCityByID(inputID uint) (bool, City)
		NewCity(inputCity City) bool
		EditCity(ID uint, newCity City) bool
		DeleteCityByID(IDS []uint) (resDel []uint)
		CallTimeDistanceTwoCities(cityNoOne, cityNoTwo City) []string
	}
)
