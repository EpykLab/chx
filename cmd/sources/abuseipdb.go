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
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/EpykLab/chx/internal/client"
	"github.com/EpykLab/chx/internal/config"
	"github.com/EpykLab/chx/internal/output"
	"github.com/charmbracelet/log"
)

const abuseipdbURL = "https://api.abuseipdb.com/api/v2/check"

// AbuseIPDBProvider queries AbuseIPDB for IP reputation.
type AbuseIPDBProvider struct {
	APIKey string
}

func (a *AbuseIPDBProvider) Name() string { return "abuseipdb" }

func (a *AbuseIPDBProvider) BuildRequest(ctx context.Context, ip string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, abuseipdbURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Key", a.APIKey)

	query := req.URL.Query()
	query.Add("ipAddress", ip)
	query.Add("maxAgeInDays", "90")
	req.URL.RawQuery = query.Encode()

	return req, nil
}

func (a *AbuseIPDBProvider) ParseResponse(body []byte) (any, error) {
	var final output.AipdbIP
	err := json.Unmarshal(body, &final)
	return &final, err
}

// GetIPInfoIpabd returns IP reputation from AbuseIPDB.
func GetIPInfoIpabd(ip string) *output.AipdbIP {
	key, err := config.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	p := &AbuseIPDBProvider{APIKey: key.ABUSE_DB_KEY}
	result, err := client.Default.Fetch(context.Background(), p, ip)
	if err != nil {
		log.Fatal(err)
	}
	return result.(*output.AipdbIP)
}
