package contact

import (
	"reflect"
	"tel-note/env"
	"tel-note/protocol"
	"testing"
)

func TestNewContact(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want []*protocol.Contact) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %p want %p", got, want)
		}
	}
	t.Run("fill new contact", func(t *testing.T) {
		contactTest := protocol.Contact{
			PersonID: 1,
			JobID:    5,
			Tel:      "",
			CellphoneCollection: []protocol.CellPhone{protocol.CellPhone{
				CellPhone:   "None",
				Description: "09121234567",
			}},
			Description: "none",
		}
		Init()
		storage.AddContact(contactTest)
		got := storage.GetContacts()
		want := env.ContactDataTest
		//slice env to one part
		want = (want)[0:1]
		//add an id to env sample
		(want)[0].Id = 1
		assertCorrectMessage(t, got, want)
	})
}
