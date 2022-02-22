package services

import (
	"tel-note/internal/services/contact"
)

func Init() {
	//init services
	var storageObject contact.StorageMemory
	contact.Storage = &storageObject
}
