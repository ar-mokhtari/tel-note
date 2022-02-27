package person

func DeletePerson(IDS []uint) []uint {
	return storage.DeletePerson(IDS)
}
