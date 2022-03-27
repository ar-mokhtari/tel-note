package callAPI

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func CallGetAPIs(URL string, Params map[string]string, Headers map[string]string) []byte {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", URL, nil)
	urlParams := url.Values{}
	for key, value := range Params {
		urlParams.Add(key, value)
	}
	for key, value := range Headers {
		req.Header.Set(key, value)
	}
	req.URL.RawQuery = urlParams.Encode()
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
	return responseData
}
