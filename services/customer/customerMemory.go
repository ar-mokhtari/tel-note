package customer

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
	"time"
)

type storageMemory protocol.CustomerStorage

func (AllCustomers storageMemory) GetCustomers() protocol.CustomerStorage {
	return protocol.CustomerStorage(AllCustomers)
}

func (AllCustomers storageMemory) NewCustomer(newCustomer protocol.Customer) {
	var LastID uint
	for ID, _ := range globalVars.CustomerMapStore {
		if ID > LastID {
			LastID = ID
		}
	}
	LastID += 1
	newCustomer.CreateAt = time.Now()
	globalVars.CustomerMapStore[LastID] = &newCustomer
	//AllCustomers.Data[LastID] = &newCustomer
}

func (AllCustomers storageMemory) EditCustomer(id uint, EditedCustomer protocol.Customer) {
	if EditedCustomer.PersonID != 0 {
		globalVars.CustomerMapStore[id].PersonID = EditedCustomer.PersonID
	}
	if EditedCustomer.Description != "" {
		globalVars.CustomerMapStore[id].Description = EditedCustomer.Description
	}
	if EditedCustomer.PersonID != 0 || EditedCustomer.Description != "" {
		globalVars.CustomerMapStore[id].UpdatedAt = time.Now()
	}
}

func (AllCustomers storageMemory) DeleteCustomerById(id uint) {
	delete(globalVars.CustomerMapStore, id)
}
