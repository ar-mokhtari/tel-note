package contact

import "net/http"

func Init() {
	var contactObject storageMemory
	storage = &contactObject
	http.HandleFunc("/get-contact", GetPool.GetContactServeHTTP)
	http.HandleFunc("/new-contact", PoolContact.NewContactServeHTTP)
	http.HandleFunc("/edit-contact", EditPool.EditContactServeHTTP)
	http.HandleFunc("/find-contact-id", FindByIDPool.FindContactByIDServeHTTP)
	http.HandleFunc("/find-contact-char", FindByCharPool.FindContactByCharServeHTTP)
	http.HandleFunc("/delete-contact-id", PoolDelContactID.DeleteContactByIDServeHTTP)
	http.HandleFunc("/delete-contact-all", PoolDelAllContact.DelAllContactServeHTTP)
}
