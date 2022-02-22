package services

import (
	"tel-note/internal/services/city"
	"tel-note/internal/services/contact"
	"tel-note/internal/services/job"
	"tel-note/internal/services/person"
	"tel-note/internal/services/sex"
)

func Init() {
	//init sex services
	var sexObject sex.StorageMemory
	sex.Storage = &sexObject

	//init city services
	var cityObject city.StorageMemory
	city.Storage = &cityObject

	//init job services
	var jobObject job.StorageMemory
	job.Storage = &jobObject

	//init person services
	var personObject person.StorageMemory
	person.Storage = &personObject

	//init contact services
	var contactObject contact.StorageMemory
	contact.Storage = &contactObject
}
