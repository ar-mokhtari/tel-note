package country

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.CountriesListR, &GetCountry)
	mux.Handle(env.CountriesCallR, &CallCountry)
	mux.Handle(env.DeleteCountryR, &DeleteCountry)
	mux.Handle(env.EditCountryR, &EditCountry)
	mux.Handle(env.FindCountryByCharR, &FindCountryChar)
	mux.Handle(env.AddCountryR, &NewCountry)
}
