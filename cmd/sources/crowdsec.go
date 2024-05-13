/*
Copyright © 2024 EpykLab contact@epyklab.com 
*/

package sources

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

const (
	crowdsecBase string = "https://cti.api.crowdsec.net/v2/"
)


func GetCrowdSecSmoke(ip string) {

	
	apiKey, ok := os.LookupEnv("crowdseckey")
	if !ok {
		fmt.Println("Crowdsec API key not found in Env")
		os.Exit(1)
	}


	uri := fmt.Sprintf("smoke/%s", ip)


	req, _ := http.NewRequest("GET", crowdsecBase+uri, nil)

	req.Header.Add("x-api-key", apiKey) 
 
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error getting information from Crowsec:", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Unable to parse response from crowsec %v", err)
	}
	
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}

	jsByte, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	output := string(jsByte)
	fmt.Println(output)


}
