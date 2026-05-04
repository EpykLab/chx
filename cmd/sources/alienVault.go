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

const avBaseURL = "https://otx.alienvault.com/api/v1/indicators/"

// AlienVaultProvider is the common base for AlienVault endpoints.
type AlienVaultProvider struct {
	APIKey string
	Path   string // e.g. "hostname/", "IPv4/", "file/"
}

func (a *AlienVaultProvider) Name() string { return "alienvault" }

func (a *AlienVaultProvider) BuildRequest(ctx context.Context, target string) (*http.Request, error) {
	url := avBaseURL + a.Path + target
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "chx")
	req.Header.Set("X-OTX-API-KEY", a.APIKey)
	return req, nil
}

// AlienVaultDomainProvider queries AlienVault for domain intelligence.
type AlienVaultDomainProvider struct{ AlienVaultProvider }

func (a *AlienVaultDomainProvider) ParseResponse(body []byte) (any, error) {
	var final output.AlienVaultDomain
	err := json.Unmarshal(body, &final)
	return &final, err
}

// AlienVaultIPProvider queries AlienVault for IP intelligence.
type AlienVaultIPProvider struct{ AlienVaultProvider }

func (a *AlienVaultIPProvider) ParseResponse(body []byte) (any, error) {
	var final output.AlientVaultIP
	err := json.Unmarshal(body, &final)
	return &final, err
}

// AlienVaultHashProvider queries AlienVault for file hash intelligence.
type AlienVaultHashProvider struct{ AlienVaultProvider }

func (a *AlienVaultHashProvider) ParseResponse(body []byte) (any, error) {
	var final output.AlienVaultHash
	err := json.Unmarshal(body, &final)
	return &final, err
}

func newAlienVaultProvider() *AlienVaultProvider {
	key, err := config.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return &AlienVaultProvider{APIKey: key.ALIEN_VAULT_KEY}
}

// GetDomainDetailsAV returns domain intelligence from AlienVault.
func GetDomainDetailsAV(domain string) *output.AlienVaultDomain {
	p := &AlienVaultDomainProvider{AlienVaultProvider: *newAlienVaultProvider()}
	p.Path = "hostname/"
	result, err := client.Default.Fetch(context.Background(), p, domain)
	if err != nil {
		log.Fatal(err)
	}
	return result.(*output.AlienVaultDomain)
}

// GetIPDetails returns IP intelligence from AlienVault.
func GetIPDetails(ip string) *output.AlientVaultIP {
	p := &AlienVaultIPProvider{AlienVaultProvider: *newAlienVaultProvider()}
	p.Path = "IPv4/"
	result, err := client.Default.Fetch(context.Background(), p, ip)
	if err != nil {
		log.Fatal(err)
	}
	return result.(*output.AlientVaultIP)
}

// GetHashDetails returns hash intelligence from AlienVault.
func GetHashDetails(hash string) *output.AlienVaultHash {
	p := &AlienVaultHashProvider{AlienVaultProvider: *newAlienVaultProvider()}
	p.Path = "file/"
	result, err := client.Default.Fetch(context.Background(), p, hash)
	if err != nil {
		log.Fatal(err)
	}
	return result.(*output.AlienVaultHash)
}
