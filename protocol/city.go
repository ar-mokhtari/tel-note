package protocol

type (
	//City []byte
	City interface {
		ID() uint
		Name() string
		EnglishName() string
		AriaCode() string
		Lat() float64
		Lng() float64
	}
	CityStorageServices interface {
		//city methods
		GetCities() []City
		//FindCityByChar(inputChar string) (status bool, res []uint) //return error
		FindCityByID(inputID uint) (error, uint) //return by ids //get:primary key find:secondary key list filter
		//NewCity(inputCity City) error
		EditCity(newCity City) error
		//DeleteCityByID(IDS []uint) (resDel []uint)
		//CallTimeDistanceTwoCities(cityNoOne, cityNoTwo City) ([]uint, bool)
	}
)
