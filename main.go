package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	abuseipdbURL string = "https://api.abuseipdb.com/api/v2/check"
)

func main() {

	var maxAgeInDays string

	if len(os.Args) < 2 {
		fmt.Println("Provide IP address as first argument\n\nEXAMPLE: ipcheck 170.205.29.2")
		os.Exit(1)
	}

	if len(os.Args) == 3 {
		maxAgeInDays = os.Args[2]
	} else {
		maxAgeInDays = "90"
	}

	ipAddr := os.Args[1]

	apiKey, ok := os.LookupEnv("abuseipdbkey")
	if !ok {
		fmt.Println("AbuseIP DB API key not found in Env")
		os.Exit(1)
	}

	req, _ := http.NewRequest("GET", abuseipdbURL, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Key", apiKey)

	query := req.URL.Query()
	query.Add("ipAddress", ipAddr)
	query.Add("maxAgeInDays", maxAgeInDays)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Unable to make request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Unable to response response: %v", err)
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}

	output, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}
