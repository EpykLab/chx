package pretty

type CrowdSecIP struct {
	AsName        string  `json:"as_name,omitempty"`
	AsNum         float64 `json:"as_num,omitempty"`
	AttackDetails []struct {
		Description string        `json:"description,omitempty"`
		Label       string        `json:"label,omitempty"`
		Name        string        `json:"name,omitempty"`
		References  []interface{} `json:"references,omitempty"`
	} `json:"attack_details,omitempty"`
	BackgroundNoise      string  `json:"background_noise,omitempty"`
	BackgroundNoiseScore float64 `json:"background_noise_score,omitempty"`
	Behaviors            []struct {
		Description string        `json:"description,omitempty"`
		Label       string        `json:"label,omitempty"`
		Name        string        `json:"name,omitempty"`
		References  []interface{} `json:"references,omitempty"`
	} `json:"behaviors,omitempty"`
	Classifications struct {
		Classifications []interface{} `json:"classifications,omitempty"`
		FalsePositives  []interface{} `json:"false_positives,omitempty"`
	} `json:"classifications,omitempty"`
	Confidence string        `json:"confidence,omitempty"`
	Cves       []interface{} `json:"cves,omitempty"`
	History    struct {
		DaysAge   float64 `json:"days_age,omitempty"`
		FirstSeen string  `json:"first_seen,omitempty"`
		FullAge   float64 `json:"full_age,omitempty"`
		LastSeen  string  `json:"last_seen,omitempty"`
	} `json:"history,omitempty"`
	Ip                  string      `json:"ip,omitempty"`
	IpRange             string      `json:"ip_range,omitempty"`
	IpRange24           interface{} `json:"ip_range_24,omitempty"`
	IpRange24Reputation interface{} `json:"ip_range_24_reputation,omitempty"`
	IpRange24Score      interface{} `json:"ip_range_24_score,omitempty"`
	IpRangeScore        float64     `json:"ip_range_score,omitempty"`
	Location            struct {
		City      interface{} `json:"city,omitempty"`
		Country   string      `json:"country,omitempty"`
		Latitude  float64     `json:"latitude,omitempty"`
		Longitude float64     `json:"longitude,omitempty"`
	} `json:"location,omitempty"`
	MitreTechniques []struct {
		Description string        `json:"description,omitempty"`
		Label       string        `json:"label,omitempty"`
		Name        string        `json:"name,omitempty"`
		References  []interface{} `json:"references,omitempty"`
	} `json:"mitre_techniques,omitempty"`
	References []struct {
		Description string        `json:"description,omitempty"`
		Label       string        `json:"label,omitempty"`
		Name        string        `json:"name,omitempty"`
		References  []interface{} `json:"references,omitempty"`
	} `json:"references,omitempty"`
	Reputation string      `json:"reputation,omitempty"`
	ReverseDns interface{} `json:"reverse_dns,omitempty"`
	Scores     struct {
		LastDay struct {
			Aggressiveness float64 `json:"aggressiveness,omitempty"`
			Anomaly        float64 `json:"anomaly,omitempty"`
			Threat         float64 `json:"threat,omitempty"`
			Total          float64 `json:"total,omitempty"`
			Trust          float64 `json:"trust,omitempty"`
		} `json:"last_day,omitempty"`
		LastMonth struct {
			Aggressiveness float64 `json:"aggressiveness,omitempty"`
			Anomaly        float64 `json:"anomaly,omitempty"`
			Threat         float64 `json:"threat,omitempty"`
			Total          float64 `json:"total,omitempty"`
			Trust          float64 `json:"trust,omitempty"`
		} `json:"last_month,omitempty"`
		LastWeek struct {
			Aggressiveness float64 `json:"aggressiveness,omitempty"`
			Anomaly        float64 `json:"anomaly,omitempty"`
			Threat         float64 `json:"threat,omitempty"`
			Total          float64 `json:"total,omitempty"`
			Trust          float64 `json:"trust,omitempty"`
		} `json:"last_week,omitempty"`
		Overall struct {
			Aggressiveness float64 `json:"aggressiveness,omitempty"`
			Anomaly        float64 `json:"anomaly,omitempty"`
			Threat         float64 `json:"threat,omitempty"`
			Total          float64 `json:"total,omitempty"`
			Trust          float64 `json:"trust,omitempty"`
		} `json:"overall,omitempty"`
	} `json:"scores,omitempty"`
	TargetCountries struct {
		BR float64
		CH float64
		CN float64
		DE float64
		FI float64
		FR float64
		GB float64
		RO float64
		TW float64
		US float64
	} `json:"target_countries,omitempty"`
}

type AipdbIP struct {
	Data struct {
		AbuseConfidenceScore float64       `json:"abuseConfidenceScore,omitempty"`
		CountryCode          string        `json:"countryCode,omitempty"`
		Domain               string        `json:"domain,omitempty"`
		Hostnames            []interface{} `json:"hostnames,omitempty"`
		IpAddress            string        `json:"ipAddress,omitempty"`
		IpVersion            float64       `json:"ipVersion,omitempty"`
		IsPublic             bool          `json:"isPublic,omitempty"`
		IsTor                bool          `json:"isTor,omitempty"`
		IsWhitelisted        bool          `json:"isWhitelisted,omitempty"`
		Isp                  string        `json:"isp,omitempty"`
		LastReportedAt       string        `json:"lastReportedAt,omitempty"`
		NumDistinctUsers     float64       `json:"numDistinctUsers,omitempty"`
		TotalReports         float64       `json:"totalReports,omitempty"`
		UsageType            string        `json:"usageType,omitempty"`
	} `json:"data,omitempty"`
}

type AlientVaultIP struct {
	AccuracyRadius float64 `json:"accuracy_radius,omitempty"`
	AreaCode       float64 `json:"area_code,omitempty"`
	Asn            string  `json:"asn,omitempty"`
	BaseIndicator  struct {
		AccessReason string  `json:"access_reason,omitempty"`
		AccessType   string  `json:"access_type,omitempty"`
		Content      string  `json:"content,omitempty"`
		Description  string  `json:"description,omitempty"`
		ID           float64 `json:"id,omitempty"`
		Indicator    string  `json:"indicator,omitempty"`
		Title        string  `json:"title,omitempty"`
		Type         string  `json:"type,omitempty"`
	} `json:"base_indicator,omitempty"`
	Charset       float64       `json:"charset,omitempty"`
	City          interface{}   `json:"city,omitempty"`
	CityData      bool          `json:"city_data,omitempty"`
	ContinentCode string        `json:"continent_code,omitempty"`
	CountryCode   string        `json:"country_code,omitempty"`
	CountryCode2  string        `json:"country_code2,omitempty"`
	CountryCode3  string        `json:"country_code3,omitempty"`
	CountryName   string        `json:"country_name,omitempty"`
	DmaCode       float64       `json:"dma_code,omitempty"`
	FalsePositive []interface{} `json:"false_positive,omitempty"`
	FlagTitle     string        `json:"flag_title,omitempty"`
	FlagURL       string        `json:"flag_url,omitempty"`
	Indicator     string        `json:"indicator,omitempty"`
	Latitude      float64       `json:"latitude,omitempty"`
	Longitude     float64       `json:"longitude,omitempty"`
	PostalCode    interface{}   `json:"postal_code,omitempty"`
	PulseInfo     struct {
		Count  float64 `json:"count,omitempty"`
		Pulses []struct {
			TLP       string
			Adversary string `json:"adversary,omitempty"`
			AttackIds []struct {
				DisplayName string `json:"display_name,omitempty"`
				ID          string `json:"id,omitempty"`
				Name        string `json:"name,omitempty"`
			} `json:"attack_ids,omitempty"`
			Author struct {
				AvatarURL    string `json:"avatar_url,omitempty"`
				ID           string `json:"id,omitempty"`
				IsFollowing  bool   `json:"is_following,omitempty"`
				IsSubscribed bool   `json:"is_subscribed,omitempty"`
				Username     string `json:"username,omitempty"`
			} `json:"author,omitempty"`
			ClonedFrom          interface{}   `json:"cloned_from,omitempty"`
			CommentCount        float64       `json:"comment_count,omitempty"`
			Created             string        `json:"created,omitempty"`
			Description         string        `json:"description,omitempty"`
			DownvotesCount      float64       `json:"downvotes_count,omitempty"`
			ExportCount         float64       `json:"export_count,omitempty"`
			FollowerCount       float64       `json:"follower_count,omitempty"`
			Groups              []interface{} `json:"groups,omitempty"`
			ID                  string        `json:"id,omitempty"`
			InGroup             bool          `json:"in_group,omitempty"`
			IndicatorCount      float64       `json:"indicator_count,omitempty"`
			IndicatorTypeCounts struct {
				CIDR float64
				IPv4 float64
				URL  float64
			} `json:"indicator_type_counts,omitempty"`
			Industries               []interface{} `json:"industries,omitempty"`
			IsAuthor                 bool          `json:"is_author,omitempty"`
			IsModified               bool          `json:"is_modified,omitempty"`
			IsSubscribing            interface{}   `json:"is_subscribing,omitempty"`
			Locked                   bool          `json:"locked,omitempty"`
			MalwareFamilies          []interface{} `json:"malware_families,omitempty"`
			Modified                 string        `json:"modified,omitempty"`
			ModifiedText             string        `json:"modified_text,omitempty"`
			Name                     string        `json:"name,omitempty"`
			Public                   float64       `json:"public,omitempty"`
			PulseSource              string        `json:"pulse_source,omitempty"`
			References               []string      `json:"references,omitempty"`
			RelatedIndicatorIsActive float64       `json:"related_indicator_is_active,omitempty"`
			RelatedIndicatorType     string        `json:"related_indicator_type,omitempty"`
			SubscriberCount          float64       `json:"subscriber_count,omitempty"`
			Tags                     []string      `json:"tags,omitempty"`
			TargetedCountries        []interface{} `json:"targeted_countries,omitempty"`
			ThreatHunterHasAgents    float64       `json:"threat_hunter_has_agents,omitempty"`
			ThreatHunterScannable    bool          `json:"threat_hunter_scannable,omitempty"`
			UpvotesCount             float64       `json:"upvotes_count,omitempty"`
			ValidatorCount           float64       `json:"validator_count,omitempty"`
			Vote                     float64       `json:"vote,omitempty"`
			VotesCount               float64       `json:"votes_count,omitempty"`
		} `json:"pulses,omitempty"`
		References []string `json:"references,omitempty"`
		Related    struct {
			Alienvault struct {
				Adversary       []interface{} `json:"adversary,omitempty"`
				Industries      []interface{} `json:"industries,omitempty"`
				MalwareFamilies []interface{} `json:"malware_families,omitempty"`
			} `json:"alienvault,omitempty"`
			Other struct {
				Adversary       []interface{} `json:"adversary,omitempty"`
				Industries      []interface{} `json:"industries,omitempty"`
				MalwareFamilies []interface{} `json:"malware_families,omitempty"`
			} `json:"other,omitempty"`
		} `json:"related,omitempty"`
	} `json:"pulse_info,omitempty"`
	Region      interface{}   `json:"region,omitempty"`
	Reputation  float64       `json:"reputation,omitempty"`
	Sections    []string      `json:"sections,omitempty"`
	Subdivision interface{}   `json:"subdivision,omitempty"`
	Type        string        `json:"type,omitempty"`
	TypeTitle   string        `json:"type_title,omitempty"`
	Validation  []interface{} `json:"validation,omitempty"`
	Whois       string        `json:"whois,omitempty"`
}
