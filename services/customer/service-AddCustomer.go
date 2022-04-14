package customer

import "tel-note/protocol"

func AddCustomer(newCustomer ...protocol.Customer) {
	for _, data := range newCustomer {
		storage.NewCustomer(data)
	}
}
