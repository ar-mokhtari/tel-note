package city

func DeleteCity(IDS []uint) []uint {
	return storage.DeleteCityByID(IDS)
}
