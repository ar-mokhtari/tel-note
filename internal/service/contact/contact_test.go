package contact

import (
	"reflect"
	"tel-note/internal/config"
	"tel-note/internal/storage"
	"testing"
)

func TestNewContact(t *testing.T) {
	MainData := &storage.AllContact{}
	assertCorrectMessage := func(t testing.TB, got, want *storage.AllContact) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %p want %p", *got, *want)
		}
	}
	t.Run("fill new contact", func(t *testing.T) {
		got := NewContact(MainData, "Alireza", "Mokhtari Garakani", "", "09121234567", "none")
		want := config.MainDataTest
		//slice config to one part
		want.ContactData = want.ContactData[6:7]
		//add an id to config sample
		(want.ContactData)[0].Id = 1
		assertCorrectMessage(t, got, &want)
	})
}
