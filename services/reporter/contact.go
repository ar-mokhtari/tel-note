package reporter

import (
	"encoding/json"
	"errors"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
	"tel-note/services/city"
	"tel-note/services/contact"
	"tel-note/services/job"
	"tel-note/services/person"
	"tel-note/services/sex"
	"time"
)

type contactReportPool struct{}

var ContactReportPool contactReportPool

type resultReport struct {
	ID           uint //Contact ID
	PID          uint //Person ID
	PersonName   string
	PersonFamily string
	DOB          time.Time
	JID          uint //Job ID
	JobName      string
	Gender       string
	Cellphone    string
	LoID         uint //Location ID of job
	JobCityName  string
	JobCity      string
	Address      string
	Description  string
}

func (allData *contactReportPool) CreateContactReport() (result []resultReport, err error) {
	allContact := contact.GetPool.GetContacts()
	for _, contact := range allContact {
		if contact.CellphoneCollection == nil {
			contact.CellphoneCollection = []protocol.CellPhone{{CellPhone: "", Description: ""}}
		}
		if err, personInfo := person.FindPersonByID(contact.PersonID); err.State {
			if err, jobInfo := job.FindJobByID(contact.JobID); err.State {
				if err, cityInfo := city.FindCityIDPool.FindCityByID(jobInfo.LocationID); err.State {
					if sexInfo, err := sex.FindSexByID(uint8(personInfo.GenderID)); err.State {
						result = append(result, resultReport{
							ID:           contact.Id,
							PID:          contact.PersonID,
							PersonName:   personInfo.FirstName,
							PersonFamily: personInfo.LastName,
							DOB:          personInfo.DOB,
							JID:          jobInfo.Id,
							JobName:      jobInfo.Name,
							Gender:       sexInfo.Name,
							Cellphone:    contact.CellphoneCollection[0].CellPhone,
							LoID:         0,
							JobCityName:  cityInfo.Name,
							Address:      contact.Address,
							Description:  contact.Description,
						})
					}
				}
			}
		} else {
			return []resultReport{}, errors.New(err.String)
		}
	}
	return result, nil
}

func (allData *contactReportPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method != env.GetMethod {
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	} else {
		if result, err := allData.CreateContactReport(); err != nil {
			json.NewEncoder(w).Encode(struct {
				State   uint
				Message string
			}{400, err.Error()})
		} else {
			json.NewEncoder(w).Encode(struct {
				State             uint
				ReportContactData []resultReport
			}{200, result})
		}
	}
}