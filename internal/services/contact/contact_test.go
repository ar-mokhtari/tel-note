package contact

//
//import (
//	"reflect"
//	"tel-note/internal/config"
//	"tel-note/internal/storage/memory"
//	"testing"
//)
//
//func TestNewContact(t *testing.T) {
//	assertCorrectMessage := func(t testing.TB, got, want *memory.AllData) {
//		t.Helper()
//		if !reflect.DeepEqual(got.ContactData, want.ContactData) {
//			t.Errorf("got %p want %p", *got, *want)
//		}
//	}
//	t.Run("fill new contact", func(t *testing.T) {
//		contactTest := memory.Contact{
//			Person: &memory.Person{
//				FirstName: "Alireza",
//				LastName:  "Mokhtari Garakani",
//			},
//			Cellphone:   "09121234567",
//			Description: "none"}
//		//got := NewContact(contactTest)
//		_, got := (*memory.AllData).AddContact(&config.MainData, contactTest)
//		want := config.MainDataTest
//		//slice config to one part
//		want.ContactData = want.ContactData[6:7]
//		//add an id to config sample
//		(want.ContactData)[0].Id = 1
//		assertCorrectMessage(t, &got, &want)
//	})
//}
