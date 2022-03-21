package contact

import "net/http"

func Init() {
	var contactObject storageMemory
	storage = &contactObject
	http.Handle("/new-contact", http.HandlerFunc(PoolContact.NewContactServeHTTP))

}
