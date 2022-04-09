package job

func DeleteJobByID(ID []uint) []uint {
	return storage.DeleteJob(ID)
}
