package country

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.CountriesListR, &GetCountry)
	http.Handle(env.CountriesCallR, &CallCountry)
}
