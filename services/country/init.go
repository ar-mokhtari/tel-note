package country

import "net/http"

func Init() {
	var countryMethods storageMemory
	storageService = &countryMethods

	http.Handle("/country-list", http.HandlerFunc(CallCountry.ServeCallCountry))
}
