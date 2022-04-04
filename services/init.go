package services

import (
	"tel-note/cli"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/general"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/reporter"
	"tel-note/services/sampleData"
	"tel-note/services/sex"
)

func Init() {
	//init general services
	general.Init()

	//init sex services
	sex.Init()

	//init city services
	city.Init()

	//init country services
	country.Init()

	//init job services
	job.Init()

	//init person services
	person.Init()

	//init customer services
	customer.Init()

	//init contact services
	contact.Init()

	//init fill data
	sampleData.Init()

	//init fill data
	reporter.Init()

	//serv handle menu
	cli.MenuHTTPServe()
}
