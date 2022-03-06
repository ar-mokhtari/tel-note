package services

import (
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/customer"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
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

	//init customer services
	customer.Init()

	//init contact services
	contact.Init()
}
