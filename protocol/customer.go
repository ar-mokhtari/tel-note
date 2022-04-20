package protocol

import "time"

type (
	Customer struct {
		PersonID    uint      `json:"person_id"`
		Description string    `json:"description"`
		CreateAt    time.Time `json:"create_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	CustomerStorage struct {
		CustomerData map[uint]*Customer
	}
	CustomerServices interface {
		NewCustomer(newCustomer Customer) error
		EditCustomer(id uint, EditedCustomer Customer) error
		DeleteCustomerById(id uint) error
		FindCustomerByID(ID uint) Customer
	}
)
