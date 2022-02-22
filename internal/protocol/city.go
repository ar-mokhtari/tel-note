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
		NewCity(inputCity City) bool
		EditCity(ID uint, newCity City) bool
		DeleteCity(IDS []uint) (resDel []uint)
	}
)
