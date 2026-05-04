package data

// Defines the data type, such as ip, hash, or domain
type Dtype string

// Defines the source of the data, such as alien vault, crowdsec, etc
type Source string

// Dtype enum
const (
	Domain Dtype = "domain"
	Hash   Dtype = "hash"
	IP     Dtype = "ip"
)

// Source enum
const (
	AlienVault Source = "alienvault"
	CrowdSec   Source = "crowdsec"
	IpAbuseDB  Source = "ipabusedb"
	VirusTotal Source = "virustotal"
)
