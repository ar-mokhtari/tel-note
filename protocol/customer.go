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
		CustomerData map[uint]*Customer
	}
	CustomerServices interface {
		NewCustomer(newCustomer Customer)
		EditCustomer(id uint, EditedCustomer Customer)
		DeleteCustomerById(id uint)
		FindCustomerByID(ID uint) Customer
	}
)
