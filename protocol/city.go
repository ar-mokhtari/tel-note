package protocol

type (
	City struct {
		Id          uint
		Name        string
		EnglishName string
		AriaCode    string
		Lat         float64
		Lng         float64
	}
	CityServices interface {
		//city methods
		GetCities() []*City
		FindCityByChar(inputChar string) (status bool, res []uint)
		FindCityByID(inputID uint) (bool, City)
		NewCity(inputCity City) bool
		EditCity(ID uint, newCity City) bool
		DeleteCityByID(IDS []uint) (resDel []uint)
		CallTimeDistanceTwoCities(cityNoOne, cityNoTwo City) ([]uint, bool)
	}
)
