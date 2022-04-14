package customer

import "tel-note/protocol"

func FindCustomerByID(ID uint) protocol.Customer {
	return storage.FindCustomerByID(ID)
}
