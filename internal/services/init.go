package services

import (
	"tel-note/internal/services/contact"
)

func Init() {
	//init contact services
	var storageObject contact.StorageMemory
	contact.Storage = &storageObject
	//init person services

}
