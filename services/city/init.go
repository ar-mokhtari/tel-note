package city

import "net/http"

func Init() {
	var cityObject storageMemory
	storage = &cityObject

	http.Handle("/distance-time-between-two-city", http.HandlerFunc(ServiceHandler.DistanceTimeServeHTTP))
}
