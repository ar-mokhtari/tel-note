package customer

import "tel-note/protocol"

func FindCustomerByID(ID uint) protocol.Customer {
	return storageCustomer.FindCustomerByID(ID)
}
