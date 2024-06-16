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
	url = "https://otx.alienvault.com/api/v1/indicators/"
)

type Headers struct {
	X_OTX_API_KEY string `json:"X-OTX-API-KEY,omitempty"`
	User_Agent    string `json:User-Agent,omitempty"`
	Content_Type  string `json:"Content-Type,omitempty"`
}

// initialize alianVault API configs
func initHeaders() *Headers {

	key, err := configs.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	apiKey := key.ALIEN_VAULT_KEY

	header := Headers{X_OTX_API_KEY: apiKey}
	header.User_Agent = "chx"
	header.Content_Type = "application/json"

	return &header
}

func makeRequest(source string, term string) {
	headers := initHeaders()

	// NOTE: this is probs a crap way of doing this. What is better
	url := fmt.Sprint(url, source, term)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Conten-Type", headers.Content_Type)
	req.Header.Add("User-Agent", headers.User_Agent)
	req.Header.Add("X-OTX-API-KEY", headers.X_OTX_API_KEY)

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

func GetDomainDetails(domain string) {
	source := "hostname/"
	makeRequest(source, domain)
}

func GetIPDetails(ip string) {
	source := "IPv4/"
	makeRequest(source, ip)
}

func GetHashDetails(hash string) {
	source := "file/"
	makeRequest(source, hash)
}
