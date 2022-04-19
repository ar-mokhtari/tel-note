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
		NewCustomer(newCustomer Customer) error
		EditCustomer(id uint, EditedCustomer Customer)
		DeleteCustomerById(id uint) error
		FindCustomerByID(ID uint) Customer
	}
)
