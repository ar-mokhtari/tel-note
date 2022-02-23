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
	sex.Init()

	//init city services
	city.Init()

	//init job services
	job.Init()

	//init person services
	person.Init()

	//init contact services
	contact.Init()
}
