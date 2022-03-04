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
		// City methods
		FindCityByChar(inputChar string) (status bool, res []uint)
		FindCityByID(inputID uint) (bool, City)
		NewCity(inputCity City) (bool, CityStorage)
		EditCity(ID uint, newCity City) (bool bool, _ CityStorage)
		DeleteCity(IDS []uint) (resDel []uint)
	}
)
