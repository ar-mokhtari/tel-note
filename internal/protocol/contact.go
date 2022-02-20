package protocol

import "context"

type Contact struct {
	Id uint `json:"id"`
	//TODO::: Person and job must be ID not struct, this way when you modify basic info, contact struct element not edit, because they store data not relation
	PersonID    uint
	JobID       uint
	Tel         string `json:"tel,omitempty"`
	Cellphone   string `json:"cellphone,omitempty"`
	Description string `json:"description,omitempty"`
}

type ContactServices interface {
	AddContact(ctx context.Context, inputContact Contact) bool
	FindContactByID(id uint) (Contact, bool)
	FindContactByChar(insertChar string) (AllData, uint)
	EditContactByID(newData Contact, ID uint) bool
	DeleteContactByID(ID uint) bool
	DeleteAll() bool
}
