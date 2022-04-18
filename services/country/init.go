package country

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.CountriesListR, &GetCountry)
	http.Handle(env.CountriesCallR, &CallCountry)
	http.Handle(env.DeleteCountryR, &DeleteCountry)
	http.Handle(env.EditCountryR, &EditCountry)
	http.Handle(env.FindCountryByCharR, &FindCountryChar)
	http.Handle(env.AddCountryR, &NewCountry)
}
