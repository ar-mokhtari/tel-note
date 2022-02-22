package contact

import (
	"reflect"
	"tel-note/internal/config"
	"tel-note/internal/protocol"
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
			JobID:       1,
			Tel:         "",
			Cellphone:   "",
			Description: "",
		}
		//TODO::: ADD PERSON AND JOB TEST FIRST (FOR PersonID and JobID)
		_, got := Storage.AddContact(contactTest)
		want := config.ContactDataTest
		//slice config to one part
		want.Data = (want.Data)[6:7]
		//add an id to config sample
		(want.Data)[0].Id = 1
		assertCorrectMessage(t, got, want)
	})
}
