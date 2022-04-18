package contact

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.ContactNewR, &NewContact)
	http.Handle(env.ContactListR, &GetContact)
	http.Handle(env.EditContactIdR, &NewContact)
	http.Handle(env.ContactFindIdR, &FindContactID)
	http.Handle(env.FindContactCharR, &FindContactChar)
	http.Handle(env.DeleteContactIdR, &DelContactID)
	http.Handle(env.DeleteAllContactsR, &DelAllContact)
}
