package protocol

type (
	CellPhone struct {
		CellPhone   string
		Description string
	}
	Contact struct {
		Id                  uint
		PersonID            uint
		JobID               uint
		Tel                 string
		CellphoneCollection []CellPhone
		Address             string
		Description         string
	}
	ContactServices interface {
		GetContacts() []*Contact
		AddContact(inputContact Contact) bool
		FindContactByID(id uint) (bool, Contact)
		FindContactByChar(insertChar string) (status bool, result []*Contact)
		EditContactByID(newData Contact, ID uint) bool
		DeleteContactByID(ID uint) bool
		DeleteAll() bool
	}
)
