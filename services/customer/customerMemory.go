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
