package contact

import "net/http"

func Init() {
	var contactObject storageMemory
	storage = &contactObject
	http.Handle("/get-contact", http.HandlerFunc(GetPool.GetContactServeHTTP))
	http.Handle("/new-contact", http.HandlerFunc(PoolContact.NewContactServeHTTP))
	http.Handle("/edit-contact", http.HandlerFunc(EditPool.EditContactServeHTTP))
}
