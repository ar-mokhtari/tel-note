package country

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.CountriesListR, &GetCountry)
	http.Handle(env.CountriesCallR, &CallCountry)
	http.Handle(env.DeleteCountryR, &DeleteCountry)
	http.Handle(env.EditCountryR, &EditCountry)
	http.Handle(env.FindCountryByCharR, &FindCountryChar)
	http.Handle(env.AddCountryR, &NewCountry)
}
