package customer

import "tel-note/protocol"

func EditCustomer(id uint, editedCustomer protocol.Customer) {
	storage.EditCustomer(id, editedCustomer)
}
