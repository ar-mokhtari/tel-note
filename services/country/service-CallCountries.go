package country

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/callAPI"
	"github.com/ar-mokhtari/tel-note/protocol"
	"github.com/ar-mokhtari/tel-note/sdk/universal"
	"net/http"
)

type callCountry struct{}

var CallCountry callCountry

func (cc *callCountry) CallCountry() []*protocol.Country {
	//generate new token
	MapTokenHeaders := map[string]string{
		"Accept":     "application/json",
		"api-token":  universal.UniversaltutorialToken,
		"user-email": universal.Email,
	}
	responseTokenData := callAPI.CallGetAPIs(
		universal.GetTokenURL,
		map[string]string{},
		MapTokenHeaders,
	)
	//var token struct {
	//	authToken string `json:"auth_token"`
	//}
	token := map[string]string{
		"auth_token": "",
	}
	json.Unmarshal(responseTokenData, &token)
	//call api with new token
	MapHeaders := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + string(token["auth_token"]),
	}
	responseData := callAPI.CallGetAPIs(
		universal.GetCountryURL,
		map[string]string{},
		MapHeaders,
	)
	var AllResult []*protocol.Country
	json.Unmarshal(responseData, &AllResult)
	return AllResult
}

func (cc *callCountry) Do() []*protocol.Country {
	return cc.CallCountry()
}
func (cc *callCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cc.Do())
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
