//go:build !mysql || !mongodb

package customer

import (
	"tel-note/protocol"
	"tel-note/services/globalVars"
	"time"
)

type (
	storageMemory protocol.CustomerStorage
	group         protocol.CustomerGroupStorage
	relation      protocol.CustomerGRelationStorage
)

var (
	storage         storageMemory
	_               protocol.CustomerServices = &storage
	storageGroup    group
	_               protocol.CustomerGroupServices = &storageGroup
	storageRelation relation
	_               protocol.CustomerGroupRelationServices = &storageRelation
)

func (allCustomers *storageMemory) NewCustomer(newCustomer protocol.Customer) {
	var LastID uint
	for ID := range globalVars.CustomerMapStore {
		if ID > LastID {
			LastID = ID
		}
	}
	LastID += 1
	newCustomer.CreateAt = time.Now()
	globalVars.CustomerMapStore[LastID] = &newCustomer
	//allCustomers.Data[LastID] = &newCustomer
}

func (allCustomers *storageMemory) EditCustomer(id uint, EditedCustomer protocol.Customer) {
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

func (allCustomers *storageMemory) DeleteCustomerById(id uint) {
	delete(globalVars.CustomerMapStore, id)
}

func (allCustomers *storageMemory) FindCustomerByID(ID uint) protocol.Customer {
	for id, customer := range globalVars.CustomerMapStore {
		if id == ID {
			return *customer
		}
	}
	return protocol.Customer{}
}

func (AllGroup *group) GetCustomerGroup() protocol.CustomerGroupStorage {
	return protocol.CustomerGroupStorage(*AllGroup)
}

func (AllGroup *group) NewGroup(groupName string) {
	var LastID uint
	for _, group := range *AllGroup {
		if group.GroupID > LastID {
			LastID = group.GroupID
		}
	}
	LastID += 1
	*AllGroup = append(*AllGroup, &protocol.CustomerGroup{
		GroupID:   LastID,
		GroupName: groupName,
	})
}

func (AllGroup *group) FindGroupByID(ID uint) protocol.CustomerGroup {
	for _, group := range *AllGroup {
		if group.GroupID == ID {
			return *group
		}
	}
	return protocol.CustomerGroup{}
}

func (AllRelation *relation) NewRelation(customerID uint, groupID uint) {
	var LastID uint
	for _, relation := range *AllRelation {
		if relation.ID > LastID {
			LastID = relation.ID
		}
	}
	LastID += 1
	*AllRelation = append(*AllRelation, &protocol.CustomerGroupRelation{
		ID:         LastID,
		CustomerID: customerID,
		GroupID:    groupID,
	})
}

func (AllRelation *relation) GetCustomerGroupRelation() protocol.CustomerGRelationStorage {
	return protocol.CustomerGRelationStorage(*AllRelation)
}

func (AllRelation *relation) FindCustomerByGroupID(ID uint) protocol.CustomerStorage {
	var result = protocol.CustomerStorage{CustomerData: make(map[uint]*protocol.Customer)}
	for _, groupRelation := range *AllRelation {
		if groupRelation.GroupID == ID {
			findCustomer := FindCustomerByID(groupRelation.CustomerID)
			result.CustomerData[groupRelation.CustomerID] = &findCustomer
		}
	}
	return result
}
