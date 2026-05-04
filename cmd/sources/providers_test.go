package sources

import (
	"context"
	"testing"

	"github.com/EpykLab/chx/internal/output"
)

func TestAbuseIPDBBuildRequest(t *testing.T) {
	p := &AbuseIPDBProvider{APIKey: "a-key"}
	req, err := p.BuildRequest(context.Background(), "1.2.3.4")
	if err != nil {
		t.Fatalf("BuildRequest() error = %v", err)
	}

	if req.Method != "GET" {
		t.Fatalf("method = %s, want GET", req.Method)
	}
	if got := req.URL.String(); got != "https://api.abuseipdb.com/api/v2/check?ipAddress=1.2.3.4&maxAgeInDays=90" {
		t.Fatalf("url = %s", got)
	}
	if got := req.Header.Get("Key"); got != "a-key" {
		t.Fatalf("Key header = %s, want a-key", got)
	}
	if got := req.Header.Get("Accept"); got != "application/json" {
		t.Fatalf("Accept header = %s, want application/json", got)
	}
}

func TestCrowdSecBuildRequest(t *testing.T) {
	p := &CrowdSecProvider{APIKey: "c-key"}
	req, err := p.BuildRequest(context.Background(), "8.8.8.8")
	if err != nil {
		t.Fatalf("BuildRequest() error = %v", err)
	}

	if got := req.URL.String(); got != "https://cti.api.crowdsec.net/v2/smoke/8.8.8.8" {
		t.Fatalf("url = %s", got)
	}
	if got := req.Header.Get("x-api-key"); got != "c-key" {
		t.Fatalf("x-api-key header = %s, want c-key", got)
	}
}

func TestVirusTotalBuildRequest(t *testing.T) {
	p := &VirusTotalProvider{APIKey: "v-key"}
	req, err := p.BuildRequest(context.Background(), "deadbeef")
	if err != nil {
		t.Fatalf("BuildRequest() error = %v", err)
	}

	if got := req.URL.String(); got != "https://www.virustotal.com/api/v3/files/deadbeef" {
		t.Fatalf("url = %s", got)
	}
	if got := req.Header.Get("x-apikey"); got != "v-key" {
		t.Fatalf("x-apikey header = %s, want v-key", got)
	}
}

func TestAlienVaultBuildRequest(t *testing.T) {
	p := &AlienVaultProvider{APIKey: "av-key", Path: "IPv4/"}
	req, err := p.BuildRequest(context.Background(), "9.9.9.9")
	if err != nil {
		t.Fatalf("BuildRequest() error = %v", err)
	}

	if got := req.URL.String(); got != "https://otx.alienvault.com/api/v1/indicators/IPv4/9.9.9.9" {
		t.Fatalf("url = %s", got)
	}
	if got := req.Header.Get("X-OTX-API-KEY"); got != "av-key" {
		t.Fatalf("X-OTX-API-KEY header = %s, want av-key", got)
	}
	if got := req.Header.Get("User-Agent"); got != "chx" {
		t.Fatalf("User-Agent = %s, want chx", got)
	}
}

func TestParseResponseTypes(t *testing.T) {
	tests := []struct {
		name string
		parse func([]byte) (any, error)
		want  any
	}{
		{name: "abuseipdb", parse: (&AbuseIPDBProvider{}).ParseResponse, want: &output.AipdbIP{}},
		{name: "crowdsec", parse: (&CrowdSecProvider{}).ParseResponse, want: &output.CrowdSecIP{}},
		{name: "virustotal", parse: (&VirusTotalProvider{}).ParseResponse, want: &output.VirusTotalHash{}},
		{name: "alienvault-domain", parse: (&AlienVaultDomainProvider{}).ParseResponse, want: &output.AlienVaultDomain{}},
		{name: "alienvault-ip", parse: (&AlienVaultIPProvider{}).ParseResponse, want: &output.AlientVaultIP{}},
		{name: "alienvault-hash", parse: (&AlienVaultHashProvider{}).ParseResponse, want: &output.AlienVaultHash{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.parse([]byte("{}"))
			if err != nil {
				t.Fatalf("ParseResponse() error = %v", err)
			}
			switch tt.want.(type) {
			case *output.AipdbIP:
				if _, ok := got.(*output.AipdbIP); !ok {
					t.Fatalf("unexpected type %T", got)
				}
			case *output.CrowdSecIP:
				if _, ok := got.(*output.CrowdSecIP); !ok {
					t.Fatalf("unexpected type %T", got)
				}
			case *output.VirusTotalHash:
				if _, ok := got.(*output.VirusTotalHash); !ok {
					t.Fatalf("unexpected type %T", got)
				}
			case *output.AlienVaultDomain:
				if _, ok := got.(*output.AlienVaultDomain); !ok {
					t.Fatalf("unexpected type %T", got)
				}
			case *output.AlientVaultIP:
				if _, ok := got.(*output.AlientVaultIP); !ok {
					t.Fatalf("unexpected type %T", got)
				}
			case *output.AlienVaultHash:
				if _, ok := got.(*output.AlienVaultHash); !ok {
					t.Fatalf("unexpected type %T", got)
				}
			}
		})
	}
}

func FuzzProviderParseResponse(f *testing.F) {
	f.Add([]byte("{}"))
	f.Add([]byte("not-json"))
	f.Add([]byte("{\"data\":{\"ipAddress\":\"1.1.1.1\"}}"))

	parsers := []func([]byte) (any, error){
		(&AbuseIPDBProvider{}).ParseResponse,
		(&CrowdSecProvider{}).ParseResponse,
		(&VirusTotalProvider{}).ParseResponse,
		(&AlienVaultDomainProvider{}).ParseResponse,
		(&AlienVaultIPProvider{}).ParseResponse,
		(&AlienVaultHashProvider{}).ParseResponse,
	}

	f.Fuzz(func(t *testing.T, body []byte) {
		for _, parse := range parsers {
			_, _ = parse(body)
		}
	})
}
