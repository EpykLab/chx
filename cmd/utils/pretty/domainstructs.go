package pretty

type AlienVaultDomain struct {
	Alexa         string `json:"alexa,omitempty"`
	BaseIndicator struct {
	} `json:"base_indicator"`
	Domain        string        `json:"domain,omitempty"`
	FalsePositive []interface{} `json:"false_positive,omitempty"`
	Indicator     string        `json:"indicator,omitempty"`
	PulseInfo     struct {
		Count      float64       `json:"count,omitempty"`
		Pulses     []interface{} `json:"pulses,omitempty"`
		References []interface{} `json:"references,omitempty"`
		Related    struct {
			Alienvault struct {
				Adversary       []interface{} `json:"adversary,omitempty"`
				Industries      []interface{} `json:"industries,omitempty"`
				MalwareFamilies []interface{} `json:"malware_families,omitempty"`
			} `json:"alienvault"`
			Other struct {
				Adversary       []interface{} `json:"adversary,omitempty"`
				Industries      []interface{} `json:"industries,omitempty"`
				MalwareFamilies []interface{} `json:"malware_families,omitempty"`
			} `json:"other"`
		} `json:"related"`
	} `json:"pulse_info"`
	Sections   []string      `json:"sections,omitempty"`
	Type       string        `json:"type,omitempty"`
	TypeTitle  string        `json:"type_title,omitempty"`
	Validation []interface{} `json:"validation,omitempty"`
	Whois      string        `json:"whois,omitempty"`
}
