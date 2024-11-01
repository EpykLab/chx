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
	"io"
	"net/http"
	"os"

	"github.com/EpykLab/chx/cmd/utils/configs"
	"github.com/EpykLab/chx/cmd/utils/pretty"
	"github.com/charmbracelet/log"
)

const (
	abuseipdbURL string = "https://api.abuseipdb.com/api/v2/check"
)

// Gets IP info from Abuse IP Database
func GetIPInfoIpabd(ip string) *pretty.AipdbIP {

	// TODO: Add in ability to adjust this time range?
	var maxAgeInDays string = "90"

	key, err := configs.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	apiKey := key.ABUSE_DB_KEY

	req, _ := http.NewRequest("GET", abuseipdbURL, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Key", apiKey)

	query := req.URL.Query()
	query.Add("ipAddress", ip)
	query.Add("maxAgeInDays", maxAgeInDays)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Unable to make request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Unable to parse response: %v", err)
	}

	var final pretty.AipdbIP

	err = json.Unmarshal(body, &final)
	if err != nil {
		log.Fatal(err)
	}

	return &final
}
