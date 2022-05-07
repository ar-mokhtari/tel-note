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
	"net/http"
)

func Init(mux *http.ServeMux) {
	//init general services
	general.Init(mux)

	//init sex services
	sex.Init(mux)

	//init city services
	city.Init(mux)

	//init country services
	country.Init(mux)

	//init job services
	job.Init(mux)

	//init person services
	person.Init(mux)

	//init customer services
	customer.Init(mux)

	//init contact services
	contact.Init(mux)

	//init fill data
	sampleData.Init(mux)

	//init fill data
	reporter.Init(mux)

	//serv handle menu
	menu.Init(mux)

	//serv www
	www.Init(mux)
}
