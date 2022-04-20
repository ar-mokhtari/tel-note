package protocol

type (
	City struct {
		Id          uint    `json:"id"`
		Name        string  `json:"name"`
		EnglishName string  `json:"english_name"`
		AriaCode    string  `json:"aria_code"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
	}
	CityServices interface {
		//city methods
		GetCities() []*City
		FindCityByChar(inputChar string) (status bool, res []uint)
		FindCityByID(inputID uint) (bool, City)
		NewCity(inputCity City) error
		EditCity(newCity City) error
		DeleteCityByID(IDS []uint) (resDel []uint)
		CallTimeDistanceTwoCities(cityNoOne, cityNoTwo City) ([]uint, bool)
	}
)
