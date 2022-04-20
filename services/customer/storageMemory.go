//go:build !mysql || !mongodb

package customer

import (
	"errors"
	"fmt"
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

func (sm *storageMemory) NewCustomer(newCustomer protocol.Customer) error {
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
	return nil
}

func (sm *storageMemory) EditCustomer(id uint, EditedCustomer protocol.Customer) error {
	if result := sm.FindCustomerByID(id); result == (protocol.Customer{}) {
		return errors.New("customer not found")
	}
	if EditedCustomer.PersonID != 0 {
		globalVars.CustomerMapStore[id].PersonID = EditedCustomer.PersonID
	}
	if EditedCustomer.Description != "" {
		globalVars.CustomerMapStore[id].Description = EditedCustomer.Description
	}
	if EditedCustomer.PersonID != 0 || EditedCustomer.Description != "" {
		globalVars.CustomerMapStore[id].UpdatedAt = time.Now()
	}
	return nil
}

func (sm *storageMemory) DeleteCustomerById(id uint) error {
	if _, ok := globalVars.CustomerMapStore[id]; !ok {
		return errors.New(fmt.Sprintf("element #%v is missing", id))
	}
	delete(globalVars.CustomerMapStore, id)
	return nil
}

func (sm *storageMemory) FindCustomerByID(ID uint) protocol.Customer {
	for id, customer := range globalVars.CustomerMapStore {
		if id == ID {
			return *customer
		}
	}
	return protocol.Customer{}
}

func (g *group) GetCustomerGroup() protocol.CustomerGroupStorage {
	return protocol.CustomerGroupStorage(*g)
}

func (g *group) NewGroup(groupName string) {
	var LastID uint
	for _, group := range *g {
		if group.GroupID > LastID {
			LastID = group.GroupID
		}
	}
	LastID += 1
	*g = append(*g, &protocol.CustomerGroup{
		GroupID:   LastID,
		GroupName: groupName,
	})
}

func (g *group) FindGroupByID(ID uint) protocol.CustomerGroup {
	for _, group := range *g {
		if group.GroupID == ID {
			return *group
		}
	}
	return protocol.CustomerGroup{}
}

func (r *relation) NewRelation(customerID uint, groupID uint) {
	var LastID uint
	for _, relation := range *r {
		if relation.ID > LastID {
			LastID = relation.ID
		}
	}
	LastID += 1
	*r = append(*r, &protocol.CustomerGroupRelation{
		ID:         LastID,
		CustomerID: customerID,
		GroupID:    groupID,
	})
}

func (r *relation) GetCustomerGroupRelation() protocol.CustomerGRelationStorage {
	return protocol.CustomerGRelationStorage(*r)
}

func (r *relation) FindCustomerGroupRelationByGroupID(ID uint) protocol.CustomerStorage {
	var result = protocol.CustomerStorage{CustomerData: make(map[uint]*protocol.Customer)}
	for _, groupRelation := range *r {
		if groupRelation.GroupID == ID {
			findCustomer := FindCustomerByID.Do(groupRelation.CustomerID)
			result.CustomerData[groupRelation.CustomerID] = &findCustomer
		}
	}
	return result
}
