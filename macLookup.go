package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
    "net/url"
	"net/http"
)

func getInputMac() string {
	
	inputMac := "44:38:39:ff:ef:57"
	
	if len(os.Args) > 1 {
		inputMac = os.Args[1]
		// fmt.Printf("Mac Address supplied: %s\n", inputMac)
	} else {
		// fmt.Printf("No Mac Address supplied, using default: %s\n", inputMac)
		inputMac = "44:38:39:ff:ef:57"
  	}

  	return inputMac
}

func buildUrl() string {

	macApiKey, found := os.LookupEnv("macApiKey")

	if !found {
		log.Fatalf("Missing macApiKey env variable")
	}

	macAddress := getInputMac()

	//fmt.Printf("Mac Address Lookup API Key: %s\n", macApiKey)

	baseUrl, err := url.Parse("https://api.macaddress.io/v1")

	if err != nil {
		log.Fatalln(err)
	}

	params := url.Values{}
	params.Add("search", macAddress)
	params.Add("apikey", macApiKey)
	params.Add("output", "vendor")

	// fmt.Printf("Querying: %s\n", baseUrl)

	baseUrl.RawQuery = params.Encode()
	
	return baseUrl.String()
}

func getMacInfo(queryUrl string) string {
	resp, err := http.Get(queryUrl)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func main() {
	// fmt.Printf("\nBeginning Execution\n\n")

	testUrl := buildUrl()

	// fmt.Printf("Query URL built: %s\n", testUrl)

	macReturn := getMacInfo(testUrl)

	// fmt.Printf("Vendor: %s\n", macReturn)
	fmt.Printf("%s", macReturn)
}