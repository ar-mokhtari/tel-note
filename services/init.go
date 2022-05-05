package services

import (
	"github.com/ar-mokhtari/tel-note/services/city"
	"github.com/ar-mokhtari/tel-note/services/contact"
	"github.com/ar-mokhtari/tel-note/services/country"
	"github.com/ar-mokhtari/tel-note/services/customer"
	"github.com/ar-mokhtari/tel-note/services/general"
	"github.com/ar-mokhtari/tel-note/services/job"
	"github.com/ar-mokhtari/tel-note/services/menu"
	"github.com/ar-mokhtari/tel-note/services/person"
	"github.com/ar-mokhtari/tel-note/services/reporter"
	"github.com/ar-mokhtari/tel-note/services/sampleData"
	"github.com/ar-mokhtari/tel-note/services/sex"
	"github.com/ar-mokhtari/tel-note/services/www"
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
	menu.Init()

	//serv www
	www.Init()
}
