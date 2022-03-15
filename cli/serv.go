package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tel-note/services/contact"
	"tel-note/services/country"
	"tel-note/services/customer"
	"tel-note/services/fillSampleData"
	"tel-note/services/general"
	"tel-note/services/person"
)

func menuHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GlobalMenu.Group)
}
func countryHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(country.CallCountry())
}
func checkIranNationalCodeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	NationalCode := r.Header.Values("NID")
	fmt.Fprintf(w, "%v", general.CheckIranNationalCode(NationalCode[0]))
}
func fillDataHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if fillSampleData.FillSimpleDataInMainData() {
		fmt.Fprintf(w, "data insert successfuly")
	} else {
		fmt.Fprintf(w, "some thing wrong")
	}
}
func getDataHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contact.GetContacts())
	json.NewEncoder(w).Encode(customer.GetCustomerGroup())
	json.NewEncoder(w).Encode(customer.GetCustomerGroupRelation())
	json.NewEncoder(w).Encode(person.GetPersons())
}
func findPersonByIDHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personID := r.FormValue("pid")
	personUID, _ := strconv.ParseUint(personID, 10, 8)
	if status, data := person.FindPersonByID(uint(personUID)); status.State {
		json.NewEncoder(w).Encode(data)
	} else {
		fmt.Fprintf(w, "not found")
	}
}
func Serv() {
	http.Handle("/menu-list", http.HandlerFunc(menuHandle))
	http.Handle("/country-list", http.HandlerFunc(countryHandle))
	http.Handle("/check-national-id", http.HandlerFunc(checkIranNationalCodeHandle))
	http.Handle("/fill-data", http.HandlerFunc(fillDataHandle))
	http.Handle("/get-data", http.HandlerFunc(getDataHandle))
	http.Handle("/find-person-id", http.HandlerFunc(findPersonByIDHandle))
	//listen for request:
	log.Fatalln(http.ListenAndServe(":1212", nil))
}
