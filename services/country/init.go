package country

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	var countryMethods storageMemory
	storageService = &countryMethods

	http.Handle(env.CountriesListR, &CallCountry)
}
