package country

func DeleteCountry(IDS []uint) uint {
	return storageService.DeleteCountry(IDS)
}
