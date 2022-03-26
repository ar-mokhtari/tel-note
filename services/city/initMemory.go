//go:build memory
// +build memory

package city

import "net/http"

func Init() {
	var cityObject storageMemory
	storage = &cityObject
	http.HandleFunc("/distance-time-between-two-city", (*DistanceTimeService).ServeHTTP(nil))
}
