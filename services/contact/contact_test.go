package contact

import (
	"gitlab.com/gocastsian/globalApp/entity"
	"gitlab.com/gocastsian/globalApp/env"
	"net/http"
	"reflect"
	"testing"
)

func TestNewContact(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want []*entity.Contact) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %p want %p", got, want)
		}
	}
	t.Run("fill new contact", func(t *testing.T) {
		contactTest := entity.Contact{
			PersonID: 1,
			JobID:    5,
			Tel:      "",
			CellphoneCollection: []entity.CellPhone{{
				CellPhone:   "0912",
				Description: "king",
			}, {CellPhone: "0912912", Description: "Home"}},
			Description: "none",
		}
		mux := *http.NewServeMux()
		Init(&mux)
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
