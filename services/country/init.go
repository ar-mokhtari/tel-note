package country

import "net/http"

func Init() {
	var countryMethods storageMemory
	storageService = &countryMethods

	http.Handle("/call-country-list", &CallCountry)
}
