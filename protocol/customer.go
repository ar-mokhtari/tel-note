package protocol

import "time"

type (
	Customer struct {
		PersonID    uint
		Description string
		CreateAt    time.Time
		UpdatedAt   time.Time
	}
	CustomerStorage struct {
		Data map[uint]*Customer
	}
	CustomerServices interface {
		GetCustomers() CustomerStorage
		NewCustomer(newCustomer Customer)
		EditCustomer(id uint, EditedCustomer Customer)
		DeleteCustomerById(id uint)
	}
)
