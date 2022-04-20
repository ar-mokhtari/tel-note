package customer

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
	"tel-note/services/globalVars"
	"time"
)

type getCustomer struct{}

var GetCustomer getCustomer

func (gc getCustomer) Do() map[uint]*protocol.Customer {
	return globalVars.CustomerMapStore
}

func (gc *getCustomer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		result := gc.Do()
		//convert map to struct
		//it's a test for behavior between map and struct
		//TODO::: after test before produce, customer in protocol define "slice of struct" instead of "map"
		type (
			Customer struct {
				CustID      uint
				PersonID    uint
				Description string
				CreateAt    time.Time
				UpdatedAt   time.Time
			}
			CustomerStorage []*Customer
		)
		var OutputData CustomerStorage
		for custID, customer := range result {
			OutputData = append(OutputData, &Customer{
				CustID:      custID,
				PersonID:    customer.PersonID,
				Description: customer.Description,
				CreateAt:    customer.CreateAt,
				UpdatedAt:   customer.UpdatedAt,
			})
		}
		json.NewEncoder(w).Encode(struct {
			State       uint
			CountResult uint
			Data        CustomerStorage
		}{200, uint(len(result)), OutputData})
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
