package protocol

type (
	Contact struct {
		Id          uint
		PersonID    uint
		JobID       uint
		Tel         string
		Cellphone   string
		Description string
	}
	ContactStorage struct {
		Data []*Contact
	}
	ContactServices interface {
		GetContacts() ContactStorage
		AddContact(inputContact Contact) bool
		FindContactByID(id uint) (bool, Contact)
		FindContactByChar(insertChar string) (status bool, result ContactStorage)
		EditContactByID(newData Contact, ID uint) bool
		DeleteContactByID(ID uint) bool
		DeleteAll() bool
	}
)
