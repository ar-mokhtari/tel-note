package customer

import "tel-note/protocol"

func GetCustomers() protocol.CustomerStorage {
	return storage.GetCustomers()
}
