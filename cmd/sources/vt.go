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

// Get hash values from Virus Total
func GetHashInfoVT(hash string) {
	headers := initVtHeaders()

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
