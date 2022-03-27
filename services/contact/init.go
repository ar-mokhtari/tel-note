package contact

import "net/http"

func Init() {
	var contactObject storageMemory
	storage = &contactObject
	http.Handle("/get-contact", &GetPool)
	http.Handle("/new-contact", &PoolContact)
	http.Handle("/edit-contact", &EditPool)
	http.Handle("/find-contact-id", &FindByIDPool)
	http.Handle("/find-contact-char", &FindByCharPool)
	http.Handle("/delete-contact-id", &PoolDelContactID)
	http.Handle("/delete-contact-all", &PoolDelAllContact)
}
