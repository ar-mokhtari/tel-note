package city

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"tel-note/SDK/neshan"
	"tel-note/env/apis"
	"tel-note/protocol"
	"time"
)

type storageMemory protocol.CityStorage

func (AllCity *storageMemory) GetCities() protocol.CityStorage {
	return protocol.CityStorage(*AllCity)
}

func (AllCity *storageMemory) FindCityByChar(inputChar string) (status bool, res []uint) {
	for _, data := range AllCity.CityData {
		if strings.Contains(data.Name, inputChar) {
			res = append(res, data.Id)
			status = true
		}
	}
	return status, res
}

func (AllCity *storageMemory) FindCityByID(inputID uint) (bool, protocol.City) {
	for _, data := range AllCity.CityData {
		if data.Id == inputID {
			return true, *data
		}
	}
	return false, protocol.City{}
}

func (AllCity *storageMemory) NewCity(inputCity protocol.City) bool {
	var LastID uint
	for _, data := range AllCity.CityData {
		if data.Id > LastID {
			LastID = data.Id
		}
	}
	LastID += 1
	result := protocol.City{
		Id:       LastID,
		Name:     inputCity.Name,
		AriaCode: inputCity.AriaCode,
		Lat:      inputCity.Lat,
		Lng:      inputCity.Lng,
	}
	AllCity.CityData = append(AllCity.CityData, &result)
	return true
}

func (AllCity *storageMemory) EditCity(ID uint, newCity protocol.City) bool {
	for index, data := range AllCity.CityData {
		if data.Id == ID {
			//TODO:: what the hell below ... is there any cleaner way for test "is it not nil?"
			if newCity.Name != "" {
				(AllCity.CityData)[index].Name = newCity.Name
			}
			if newCity.Name != "" {
				(AllCity.CityData)[index].Name = newCity.Name
			}
			if newCity.AriaCode != "" {
				(AllCity.CityData)[index].AriaCode = newCity.AriaCode
			}
			return true
		}
	}
	return false
}

func (AllCity *storageMemory) DeleteCityByID(IDS []uint) (resDel []uint) {
	for _, id := range IDS {
		for index, data := range AllCity.CityData {
			if data.Id == id {
				AllCity.CityData = append((AllCity.CityData)[:index], (AllCity.CityData)[index+1:]...)
				resDel = append(resDel, id)
			}
		}
	}
	return resDel
}

func (AllCity *storageMemory) CallTimeDistanceTwoCities(cityNoOne, cityNoTwo protocol.City) ([]string, bool) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	urlParams := url.Values{}
	urlParams.Add("type", "car")
	urlParams.Add("origins", fmt.Sprintf("%f,%f", cityNoOne.Lat, cityNoOne.Lng))
	urlParams.Add("destinations", fmt.Sprintf("%f,%f", cityNoTwo.Lat, cityNoTwo.Lng))
	req, _ := http.NewRequest("GET", apis.NeshanURL, nil)
	req.URL.RawQuery = urlParams.Encode()
	req.Header.Set("Api-Key", apis.NeshanAPIKey)
	response, callErr := client.Do(req)
	if callErr != nil {
		fmt.Println(callErr.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalln(readErr)
	}
	var AllResult neshan.DistanceMatrix
	json.Unmarshal(responseData, &AllResult)
	switch AllResult.Rows == nil {
	case true:
		return []string{}, false
	case false:
		return []string{
			AllResult.Rows[0].Elements[0].Duration.Text,
			AllResult.Rows[0].Elements[0].Distance.Text,
		}, true
	}
	return []string{}, false
}
