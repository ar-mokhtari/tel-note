package contact

import (
	"reflect"
	"tel-note/env"
	"tel-note/protocol"
	"testing"
)

func TestNewContact(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want protocol.ContactStorage) {
		t.Helper()
		if !reflect.DeepEqual(got.Data, want.Data) {
			t.Errorf("got %p want %p", got, want)
		}
	}
	t.Run("fill new contact", func(t *testing.T) {
		contactTest := protocol.Contact{
			PersonID:    1,
			JobID:       5,
			Tel:         "",
			Cellphone:   "09121234567",
			Description: "none",
		}
		Init()
		_, got := storage.AddContact(contactTest)
		want := env.ContactDataTest
		//slice env to one part
		want.Data = (want.Data)[0:1]
		//add an id to env sample
		(want.Data)[0].Id = 1
		assertCorrectMessage(t, got, want)
	})
}
