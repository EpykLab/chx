/*
Copyright © 2024 contact@epyklab.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package sources

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/EpykLab/chx/cmd/utils/configs"
	"github.com/EpykLab/chx/cmd/utils/pretty"
	"github.com/charmbracelet/log"
)

const (
	// AlienVaultURL is the URL for the AlienVault API
	url = "https://otx.alienvault.com/api/v1/indicators/"
)

type Headers struct {
	X_OTX_API_KEY string `json:"X-OTX-API-KEY,omitempty"`
	User_Agent    string `json:"User-Agent,omitempty"`
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

func makeRequest(source string, term string) []byte {
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

	return body
}

// Get details about domain from AlienVault
func GetDomainDetailsAV(domain string) *pretty.AlienVaultDomain {

	source := "hostname/"
	result := makeRequest(source, domain)

	var final pretty.AlienVaultDomain

	err := json.Unmarshal(result, &final)
	if err != nil {
		log.Fatal(err)
	}

	return &final
}

// Get details about IP from AlienVault
func GetIPDetails(ip string) *pretty.AlientVaultIP {
	source := "IPv4/"
	result := makeRequest(source, ip)

	var final pretty.AlientVaultIP

	err := json.Unmarshal(result, &final)
	if err != nil {
		log.Fatal(err)
	}

	return &final
}

// Get details about hash from AlienVault
func GetHashDetails(hash string) *pretty.AlienVaultHash {
	source := "file/"
	result := makeRequest(source, hash)

	var final pretty.AlienVaultHash

	err := json.Unmarshal(result, &final)
	if err != nil {
		log.Fatal(err)
	}

	return &final
}
