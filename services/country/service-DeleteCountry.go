package country

func DeleteCountry(IDS []uint) uint {
	return storage.DeleteCountry(IDS)
}
