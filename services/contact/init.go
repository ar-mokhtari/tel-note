package contact

import "net/http"

func Init() {
	var contactObject storageMemory
	storage = &contactObject
	http.Handle("/get-contact", http.HandlerFunc(GetPool.GetContactServeHTTP))
	http.Handle("/new-contact", http.HandlerFunc(PoolContact.NewContactServeHTTP))
	http.Handle("/edit-contact", http.HandlerFunc(EditPool.EditContactServeHTTP))
	http.Handle("/find-contact-id", http.HandlerFunc(FindByIDPool.FindContactByIDServeHTTP))
	http.Handle("/find-contact-char", http.HandlerFunc(FindByCharPool.FindContactByCharServeHTTP))
	http.Handle("/delete-contact-id", http.HandlerFunc(PoolDelContactID.DeleteContactByIDServeHTTP))
	http.Handle("/delete-contact-all", http.HandlerFunc(PoolDelAllContact.DelAllContactServeHTTP))
}
