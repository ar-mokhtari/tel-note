package city

import "net/http"

func Init() {

	http.Handle("/distance-time-between-two-city", &DistanceTimeService)
	http.Handle("/find-city-char", &FindByCharService)
}
