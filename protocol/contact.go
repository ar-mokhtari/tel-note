package protocol

type (
	CellPhone struct {
		CellPhone   string
		Description string
	}
	Contact struct {
		Id                  uint        `json:"id"`
		PersonID            uint        `json:"person_id"`
		JobID               uint        `json:"job_id"`
		Tel                 string      `json:"tel"`
		CellphoneCollection []CellPhone `json:"cellphone_collection"`
		Address             string      `json:"address"`
		Description         string      `json:"description"`
	}
	ContactServices interface {
		GetContacts() []*Contact
		AddContact(inputContact Contact) error
		FindContactByID(id uint) (bool, Contact)
		FindContactByChar(insertChar string) (status bool, result []*Contact)
		EditContactByID(newData Contact) error
		DeleteContactByID(ID uint) bool
		DeleteAll() bool
	}
)
