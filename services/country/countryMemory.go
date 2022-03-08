package country

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"tel-note/env/apis"
	"tel-note/protocol"
)

type storageCountry protocol.CityStorage

func (AllCountries storageCountry) GetCountry() protocol.CountryStorage {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apis.UniversalTutorialURL, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", apis.UniversalTutorialAPIKey)
	response, callErr := client.Do(req)
	if callErr != nil {
		fmt.Println(callErr.Error())
		os.Exit(1)
	}
	responseData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalln(readErr)
	}
	fmt.Println(string(responseData))
	return protocol.CountryStorage{}
}
