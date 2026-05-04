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
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/EpykLab/chx/internal/client"
	"github.com/EpykLab/chx/internal/config"
	"github.com/EpykLab/chx/internal/output"
	"github.com/charmbracelet/log"
)

const crowdsecBase = "https://cti.api.crowdsec.net/v2/"

// CrowdSecProvider queries CrowdSec for IP intelligence.
type CrowdSecProvider struct {
	APIKey string
}

func (c *CrowdSecProvider) Name() string { return "crowdsec" }

func (c *CrowdSecProvider) BuildRequest(ctx context.Context, ip string) (*http.Request, error) {
	uri := fmt.Sprintf("smoke/%s", ip)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, crowdsecBase+uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.APIKey)
	return req, nil
}

func (c *CrowdSecProvider) ParseResponse(body []byte) (any, error) {
	var final output.CrowdSecIP
	err := json.Unmarshal(body, &final)
	return &final, err
}

// GetCrowdSecSmoke returns IP intelligence from CrowdSec.
func GetCrowdSecSmoke(ip string) *output.CrowdSecIP {
	key, err := config.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	p := &CrowdSecProvider{APIKey: key.CROWD_SEC_KEY}
	result, err := client.Default.Fetch(context.Background(), p, ip)
	if err != nil {
		if strings.Contains(err.Error(), "not-found error 404") {
			return &output.CrowdSecIP{Ip: ip}
		}
		log.Fatal(err)
	}
	return result.(*output.CrowdSecIP)
}
