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

	"github.com/EpykLab/chx/internal/client"
	"github.com/EpykLab/chx/internal/config"
	"github.com/EpykLab/chx/internal/output"
	"github.com/charmbracelet/log"
)

const vtUrl = "https://www.virustotal.com/api/v3/files/"

// VirusTotalProvider queries VirusTotal for file hash intelligence.
type VirusTotalProvider struct {
	APIKey string
}

func (v *VirusTotalProvider) Name() string { return "virustotal" }

func (v *VirusTotalProvider) BuildRequest(ctx context.Context, hash string) (*http.Request, error) {
	url := fmt.Sprint(vtUrl, hash)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "chx")
	req.Header.Set("x-apikey", v.APIKey)
	return req, nil
}

func (v *VirusTotalProvider) ParseResponse(body []byte) (any, error) {
	var final output.VirusTotalHash
	err := json.Unmarshal(body, &final)
	return &final, err
}

// GetHashInfoVT returns hash intelligence from VirusTotal.
func GetHashInfoVT(hash string) *output.VirusTotalHash {
	key, err := config.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	p := &VirusTotalProvider{APIKey: key.VIRUS_TOTAL_KEY}
	result, err := client.Default.Fetch(context.Background(), p, hash)
	if err != nil {
		log.Fatal(err)
	}
	return result.(*output.VirusTotalHash)
}
