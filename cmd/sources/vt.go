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
	// AlienVaultURL is the URL for the AlienVault API
	vtUrl = "https://www.virustotal.com/api/v3/files/"
)

type VtHeaders struct {
	X_APIKEY     string `json:"x-apikey,omitempty"`
	User_Agent   string `json:"User-Agent,omitempty"`
	Content_Type string `json:"Content-Type,omitempty"`
}

// initialize alianVault API configs
func initVtHeaders() *VtHeaders {

	key, err := configs.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	apiKey := key.VIRUS_TOTAL_KEY

	header := VtHeaders{X_APIKEY: apiKey}
	header.User_Agent = "chx"
	header.Content_Type = "application/json"

	return &header
}

func MakeVtRequest(hash string) {
	headers := initVtHeaders()

	// NOTE: this is probs a crap way of doing this. What is better
	url := fmt.Sprint(vtUrl, hash)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Conten-Type", headers.Content_Type)
	req.Header.Add("User-Agent", headers.User_Agent)
	req.Header.Add("x-apikey", headers.X_APIKEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
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
