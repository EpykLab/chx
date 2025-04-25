package pretty

type AlienVaultHash struct {
	BaseIndicator struct {
		AccessReason string  `json:"access_reason,omitempty"`
		AccessType   string  `json:"access_type,omitempty"`
		Content      string  `json:"content,omitempty"`
		Description  string  `json:"description,omitempty"`
		ID           float64 `json:"id,omitempty"`
		Indicator    string  `json:"indicator,omitempty"`
		Title        string  `json:"title,omitempty"`
		Type         string  `json:"type,omitempty"`
	} `json:"base_indicator"`
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
}

type VirusTotalHash struct {
	Data struct {
		Attributes struct {
			Authentihash           string  `json:"authentihash,omitempty"`
			CreationDate           float64 `json:"creation_date,omitempty"`
			CrowdsourcedIdsResults []struct {
				AlertContext []struct {
					DestIp   string   `json:"dest_ip,omitempty"`
					DestPort float64  `json:"dest_port,omitempty"`
					Ja3      []string `json:"ja3,omitempty"`
					Ja3s     []string `json:"ja3s,omitempty"`
				} `json:"alert_context,omitempty"`
				AlertSeverity  string   `json:"alert_severity,omitempty"`
				RuleID         string   `json:"rule_id,omitempty"`
				RuleMsg        string   `json:"rule_msg,omitempty"`
				RuleRaw        string   `json:"rule_raw,omitempty"`
				RuleReferences []string `json:"rule_references,omitempty"`
				RuleSource     string   `json:"rule_source,omitempty"`
				RuleURL        string   `json:"rule_url,omitempty"`
			} `json:"crowdsourced_ids_results,omitempty"`
			CrowdsourcedIdsStats struct {
				High   float64 `json:"high,omitempty"`
				Info   float64 `json:"info,omitempty"`
				Low    float64 `json:"low,omitempty"`
				Medium float64 `json:"medium,omitempty"`
			} `json:"crowdsourced_ids_stats"`
			Detectiteasy struct {
				Filetype string `json:"filetype,omitempty"`
				Values   []struct {
					Info    string `json:"info,omitempty"`
					Name    string `json:"name,omitempty"`
					Type    string `json:"type,omitempty"`
					Version string `json:"version,omitempty"`
				} `json:"values,omitempty"`
			} `json:"detectiteasy"`
			FirstSubmissionDate float64 `json:"first_submission_date,omitempty"`
			LastAnalysisDate    float64 `json:"last_analysis_date,omitempty"`
			LastAnalysisResults struct {
				ALYac struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ALYac"`
				APEX struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"APEX"`
				AVG struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"AVG"`
				Acronis struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Acronis"`
				AhnLab_V3 struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"AhnLab_V3"`
				Alibaba struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Alibaba"`
				Antiy_AVL struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Antiy-AVL"`
				Arcabit struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Arcabit"`
				Avast struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Avast"`
				Avast_Mobile struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Avast-Mobile"`
				Avira struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Avira"`
				Baidu struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Baidu"`
				BitDefender struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				}
				BitDefenderFalx struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"BitDefenderFalx"`
				Bkav struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Bkav"`
				CAT_QuickHeal struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"CAT_QuickHeal"`
				CMC struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"CMC"`
				CTX struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				}
				ClamAV struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ClamAV"`
				CrowdStrike struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"CrowdStrike"`
				Cybereason struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Cybereason"`
				Cylance struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Cylance"`
				Cynet struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Cynet"`
				DeepInstinct struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"DeepInstinct"`
				DrWeb struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"DrWeb"`
				ESET_NOD32 struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ESET-NOD32"`
				Elastic struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Elastic"`
				Emsisoft struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Emsisoft"`
				F_Secure struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"F-Secure"`
				FireEye struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"FireEye"`
				Fortinet struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Fortinet"`
				GData struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"GData"`
				Google struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Google"`
				Gridinsoft struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Gridinsoft"`
				Ikarus struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Ikarus"`
				Jiangmin struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Jiangmin"`
				K7AntiVirus struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"K7AntiVirus"`
				K7GW struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"K7GW"`
				Kaspersky struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Kaspersky"`
				Kingsoft struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Kingsoft"`
				Lionic struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Lionic"`
				Malwarebytes struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Malwarebytes"`
				MaxSecure struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"MaxSecure"`
				McAfee struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"McAfee"`
				McAfeeD struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"McAfeeD"`
				MicroWorld_EScan struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"MicroWorld-eScan"`
				Microsoft struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Microsoft"`
				NANO_Antivirus struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"NANO-Antivirus"`
				Paloalto struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Paloalto"`
				Panda struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Panda"`
				Rising struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Rising"`
				SUPERAntiSpyware struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"SUPERAntiSpyware"`
				Sangfor struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Sangfor"`
				SentinelOne struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"SentinelOne"`
				Skyhigh struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Skyhigh"`
				Sophos struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Sophos"`
				Symantec struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Symantec"`
				SymantecMobileInsight struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"SymantecMobileInsight"`
				TACHYON struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"TACHYON"`
				Tencent struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Tencent"`
				Trapmine struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Trapmine"`
				TrendMicro struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"TrendMicro"`
				TrendMicro_HouseCall struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"TrendMicro-HouseCall"`
				Trustlook struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Trustlook"`
				VBA32 struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"VBA32"`
				VIPRE struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"VIPRE"`
				Varist struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Varist"`
				ViRobot struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ViRobot"`
				VirIT struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"VirIT"`
				Webroot struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Webroot"`
				Xcitium struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Xcitium"`
				Yandex struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Yandex"`
				Zillya struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Zillya"`
				ZoneAlarm struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ZoneAlarm"`
				Zoner struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Zoner"`
				Alibabacloud struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"alibabacloud"`
				Huorong struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"huorong"`
				Tehtris struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"tehtris"`
			} `json:"last_analysis_results"`
			LastAnalysisStats struct {
				Confirmed_Timeout float64 `json:"confirmed-timeout,omitempty"`
				Failure           float64 `json:"failure,omitempty"`
				Harmless          float64 `json:"harmless,omitempty"`
				Malicious         float64 `json:"malicious,omitempty"`
				Suspicious        float64 `json:"suspicious,omitempty"`
				Timeout           float64 `json:"timeout,omitempty"`
				Type_Unsupported  float64 `json:"type-unsupported,omitempty"`
				Undetected        float64 `json:"undetected,omitempty"`
			} `json:"last_analysis_stats"`
			LastModificationDate float64  `json:"last_modification_date,omitempty"`
			LastSubmissionDate   float64  `json:"last_submission_date,omitempty"`
			Magic                string   `json:"magic,omitempty"`
			Magika               string   `json:"magika,omitempty"`
			Md5                  string   `json:"md5,omitempty"`
			MeaningfulName       string   `json:"meaningful_name,omitempty"`
			Names                []string `json:"names,omitempty"`
			PeInfo               struct {
				CompilerProductVersions []string `json:"compiler_product_versions,omitempty"`
				EntryPoint              float64  `json:"entry_point,omitempty"`
				Exports                 []string `json:"exports,omitempty"`
				Imphash                 string   `json:"imphash,omitempty"`
				ImportList              []struct {
					ImportedFunctions []string `json:"imported_functions,omitempty"`
					LibraryName       string   `json:"library_name,omitempty"`
				} `json:"import_list,omitempty"`
				MachineType     float64 `json:"machine_type,omitempty"`
				ResourceDetails []struct {
					Chi2     float64 `json:"chi2,omitempty"`
					Entropy  float64 `json:"entropy,omitempty"`
					Filetype string  `json:"filetype,omitempty"`
					Lang     string  `json:"lang,omitempty"`
					Sha256   string  `json:"sha256,omitempty"`
					Type     string  `json:"type,omitempty"`
				} `json:"resource_details,omitempty"`
				ResourceLangs struct {
					ENGLISH_US float64 `json:"ENGLISH US,omitempty"`
				} `json:"resource_langs"`
				ResourceTypes struct {
					RTMANIFEST float64 `json:"RT_MANIFEST,omitempty"`
				} `json:"resource_types"`
				RichPeHeaderHash string `json:"rich_pe_header_hash,omitempty"`
				Sections         []struct {
					Chi2           float64 `json:"chi2,omitempty"`
					Entropy        float64 `json:"entropy,omitempty"`
					Flags          string  `json:"flags,omitempty"`
					Md5            string  `json:"md5,omitempty"`
					Name           string  `json:"name,omitempty"`
					RawSize        float64 `json:"raw_size,omitempty"`
					VirtualAddress float64 `json:"virtual_address,omitempty"`
					VirtualSize    float64 `json:"virtual_size,omitempty"`
				} `json:"sections,omitempty"`
				Timestamp float64 `json:"timestamp,omitempty"`
			} `json:"pe_info"`
			PopularThreatClassification struct {
				PopularThreatCategory []struct {
					Count float64 `json:"count,omitempty"`
					Value string  `json:"value,omitempty"`
				} `json:"popular_threat_category,omitempty"`
				PopularThreatName []struct {
					Count float64 `json:"count,omitempty"`
					Value string  `json:"value,omitempty"`
				} `json:"popular_threat_name,omitempty"`
				SuggestedThreatLabel string `json:"suggested_threat_label,omitempty"`
			} `json:"popular_threat_classification"`
			Reputation      float64 `json:"reputation,omitempty"`
			SandboxVerdicts struct {
				C2AE struct {
					Category              string   `json:"category,omitempty"`
					MalwareClassification []string `json:"malware_classification,omitempty"`
					SandboxName           string   `json:"sandbox_name,omitempty"`
				}
			} `json:"sandbox_verdicts"`
			Sha1                 string `json:"sha1,omitempty"`
			Sha256               string `json:"sha256,omitempty"`
			SigmaAnalysisResults []struct {
				MatchContext []struct {
					Values struct {
						EventID           string
						PrivilegeList     string
						SubjectDomainName string
						SubjectLogonId    string
						SubjectUserName   string
						SubjectUserSid    string
					} `json:"values"`
				} `json:"match_context,omitempty"`
				RuleAuthor      string `json:"rule_author,omitempty"`
				RuleDescription string `json:"rule_description,omitempty"`
				RuleID          string `json:"rule_id,omitempty"`
				RuleLevel       string `json:"rule_level,omitempty"`
				RuleSource      string `json:"rule_source,omitempty"`
				RuleTitle       string `json:"rule_title,omitempty"`
			} `json:"sigma_analysis_results,omitempty"`
			SigmaAnalysisStats struct {
				Critical float64 `json:"critical,omitempty"`
				High     float64 `json:"high,omitempty"`
				Low      float64 `json:"low,omitempty"`
				Medium   float64 `json:"medium,omitempty"`
			} `json:"sigma_analysis_stats"`
			SigmaAnalysisSummary struct {
				Sigma_Integrated_Rule_Set__GitHub_ struct {
					Critical float64 `json:"critical,omitempty"`
					High     float64 `json:"high,omitempty"`
					Low      float64 `json:"low,omitempty"`
					Medium   float64 `json:"medium,omitempty"`
				} `json:"Sigma Integrated Rule Set (GitHub)"`
			} `json:"sigma_analysis_summary"`
			Size           float64  `json:"size,omitempty"`
			Ssdeep         string   `json:"ssdeep,omitempty"`
			Tags           []string `json:"tags,omitempty"`
			TimesSubmitted float64  `json:"times_submitted,omitempty"`
			Tlsh           string   `json:"tlsh,omitempty"`
			TotalVotes     struct {
				Harmless  float64 `json:"harmless,omitempty"`
				Malicious float64 `json:"malicious,omitempty"`
			} `json:"total_votes"`
			Trid []struct {
				FileType    string  `json:"file_type,omitempty"`
				Probability float64 `json:"probability,omitempty"`
			} `json:"trid,omitempty"`
			TypeDescription string   `json:"type_description,omitempty"`
			TypeExtension   string   `json:"type_extension,omitempty"`
			TypeTag         string   `json:"type_tag,omitempty"`
			TypeTags        []string `json:"type_tags,omitempty"`
			UniqueSources   float64  `json:"unique_sources,omitempty"`
			Vhash           string   `json:"vhash,omitempty"`
		} `json:"attributes"`
		ID    string `json:"id,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links"`
		Type string `json:"type,omitempty"`
	} `json:"data"`
}
