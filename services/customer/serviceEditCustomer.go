package customer

import "tel-note/protocol"

func EditCustomer(id uint, editedCustomer protocol.Customer) {
	storageCustomer.EditCustomer(id, editedCustomer)
}
