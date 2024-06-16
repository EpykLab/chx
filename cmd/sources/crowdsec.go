/*
Copyright © 2024 EpykLab contact@epyklab.com
*/

package sources

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/EpykLab/chx/cmd/utils/configs"
	"github.com/charmbracelet/log"
)

const (
	crowdsecBase string = "https://cti.api.crowdsec.net/v2/"
)

func GetCrowdSecSmoke(ip string) {

	key, err := configs.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	apiKey := key.CROWD_SEC_KEY

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
