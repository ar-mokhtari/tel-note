package customer

import "tel-note/protocol"

func GetCustomerGroup() protocol.CustomerGroupStorage {
	return storageGroup.GetCustomerGroup()
}
