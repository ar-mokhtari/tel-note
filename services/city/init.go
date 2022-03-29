package city

import "net/http"

func Init() {
	http.Handle("/distance-time-between-two-city", &DistanceTimeService)
	http.Handle("/find-city-char", &FindByCharService)
	http.Handle("/delete-city", &DeleteCityPool)
	http.Handle("/edit-city", &EditCityPool)
	http.Handle("/find-city-id", &FindCityIDPool)
	http.Handle("/get-city", &GetCityPool)
	http.Handle("/add-city", &NewCityPool)
}
