package contact

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.ContactNewR, &NewContact)
	mux.Handle(env.ContactListR, &GetContact)
	mux.Handle(env.EditContactIdR, &NewContact)
	mux.Handle(env.ContactFindIdR, &FindContactID)
	mux.Handle(env.FindContactCharR, &FindContactChar)
	mux.Handle(env.DeleteContactIdR, &DelContactID)
	mux.Handle(env.DeleteAllContactsR, &DelAllContact)
}
