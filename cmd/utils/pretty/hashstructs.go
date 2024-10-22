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
	} `json:"base_indicator,omitempty"`
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
			} `json:"alienvault,omitempty"`
			Other struct {
				Adversary       []interface{} `json:"adversary,omitempty"`
				Industries      []interface{} `json:"industries,omitempty"`
				MalwareFamilies []interface{} `json:"malware_families,omitempty"`
			} `json:"other,omitempty"`
		} `json:"related,omitempty"`
	} `json:"pulse_info,omitempty"`
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
			} `json:"crowdsourced_ids_stats,omitempty"`
			Detectiteasy struct {
				Filetype string `json:"filetype,omitempty"`
				Values   []struct {
					Info    string `json:"info,omitempty"`
					Name    string `json:"name,omitempty"`
					Type    string `json:"type,omitempty"`
					Version string `json:"version,omitempty"`
				} `json:"values,omitempty"`
			} `json:"detectiteasy,omitempty"`
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
				} `json:"ALYac,omitempty"`
				APEX struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"APEX,omitempty"`
				AVG struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"AVG,omitempty"`
				Acronis struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Acronis,omitempty"`
				AhnLab_V3 struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"AhnLab_V3,omitempty"`
				Alibaba struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Alibaba,omitempty"`
				Antiy_AVL struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Antiy-AVL,omitempty"`
				Arcabit struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Arcabit,omitempty"`
				Avast struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Avast,omitempty"`
				Avast_Mobile struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Avast-Mobile,omitempty"`
				Avira struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Avira,omitempty"`
				Baidu struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Baidu,omitempty"`
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
				} `json:"BitDefenderFalx,omitempty"`
				Bkav struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Bkav,omitempty"`
				CAT_QuickHeal struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"CAT_QuickHeal,omitempty"`
				CMC struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"CMC,omitempty"`
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
				} `json:"ClamAV,omitempty"`
				CrowdStrike struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"CrowdStrike,omitempty"`
				Cybereason struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Cybereason,omitempty"`
				Cylance struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Cylance,omitempty"`
				Cynet struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Cynet,omitempty"`
				DeepInstinct struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"DeepInstinct,omitempty"`
				DrWeb struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"DrWeb,omitempty"`
				ESET_NOD32 struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ESET-NOD32,omitempty"`
				Elastic struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Elastic,omitempty"`
				Emsisoft struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Emsisoft,omitempty"`
				F_Secure struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"F-Secure,omitempty"`
				FireEye struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"FireEye,omitempty"`
				Fortinet struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Fortinet,omitempty"`
				GData struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"GData,omitempty"`
				Google struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Google,omitempty"`
				Gridinsoft struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Gridinsoft,omitempty"`
				Ikarus struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Ikarus,omitempty"`
				Jiangmin struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Jiangmin,omitempty"`
				K7AntiVirus struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"K7AntiVirus,omitempty"`
				K7GW struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"K7GW,omitempty"`
				Kaspersky struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Kaspersky,omitempty"`
				Kingsoft struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Kingsoft,omitempty"`
				Lionic struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Lionic,omitempty"`
				Malwarebytes struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Malwarebytes,omitempty"`
				MaxSecure struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"MaxSecure,omitempty"`
				McAfee struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"McAfee,omitempty"`
				McAfeeD struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"McAfeeD,omitempty"`
				MicroWorld_EScan struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"MicroWorld-eScan,omitempty"`
				Microsoft struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Microsoft,omitempty"`
				NANO_Antivirus struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"NANO-Antivirus,omitempty"`
				Paloalto struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Paloalto,omitempty"`
				Panda struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Panda,omitempty"`
				Rising struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Rising,omitempty"`
				SUPERAntiSpyware struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"SUPERAntiSpyware,omitempty"`
				Sangfor struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Sangfor,omitempty"`
				SentinelOne struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"SentinelOne,omitempty"`
				Skyhigh struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Skyhigh,omitempty"`
				Sophos struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Sophos,omitempty"`
				Symantec struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Symantec,omitempty"`
				SymantecMobileInsight struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"SymantecMobileInsight,omitempty"`
				TACHYON struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"TACHYON,omitempty"`
				Tencent struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Tencent,omitempty"`
				Trapmine struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Trapmine,omitempty"`
				TrendMicro struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"TrendMicro,omitempty"`
				TrendMicro_HouseCall struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"TrendMicro-HouseCall,omitempty"`
				Trustlook struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Trustlook,omitempty"`
				VBA32 struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"VBA32,omitempty"`
				VIPRE struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"VIPRE,omitempty"`
				Varist struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Varist,omitempty"`
				ViRobot struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ViRobot,omitempty"`
				VirIT struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"VirIT,omitempty"`
				Webroot struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Webroot,omitempty"`
				Xcitium struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Xcitium,omitempty"`
				Yandex struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Yandex,omitempty"`
				Zillya struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"Zillya,omitempty"`
				ZoneAlarm struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"ZoneAlarm,omitempty"`
				Zoner struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"Zoner,omitempty"`
				Alibabacloud struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"alibabacloud,omitempty"`
				Huorong struct {
					Category      string `json:"category,omitempty"`
					EngineName    string `json:"engine_name,omitempty"`
					EngineUpdate  string `json:"engine_update,omitempty"`
					EngineVersion string `json:"engine_version,omitempty"`
					Method        string `json:"method,omitempty"`
					Result        string `json:"result,omitempty"`
				} `json:"huorong,omitempty"`
				Tehtris struct {
					Category      string      `json:"category,omitempty"`
					EngineName    string      `json:"engine_name,omitempty"`
					EngineUpdate  string      `json:"engine_update,omitempty"`
					EngineVersion string      `json:"engine_version,omitempty"`
					Method        string      `json:"method,omitempty"`
					Result        interface{} `json:"result,omitempty"`
				} `json:"tehtris,omitempty"`
			} `json:"last_analysis_results,omitempty"`
			LastAnalysisStats struct {
				Confirmed_Timeout float64 `json:"confirmed-timeout,omitempty"`
				Failure           float64 `json:"failure,omitempty"`
				Harmless          float64 `json:"harmless,omitempty"`
				Malicious         float64 `json:"malicious,omitempty"`
				Suspicious        float64 `json:"suspicious,omitempty"`
				Timeout           float64 `json:"timeout,omitempty"`
				Type_Unsupported  float64 `json:"type-unsupported,omitempty"`
				Undetected        float64 `json:"undetected,omitempty"`
			} `json:"last_analysis_stats,omitempty"`
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
				} `json:"resource_langs,omitempty"`
				ResourceTypes struct {
					RTMANIFEST float64 `json:"RT_MANIFEST,omitempty"`
				} `json:"resource_types,omitempty"`
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
			} `json:"pe_info,omitempty"`
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
			} `json:"popular_threat_classification,omitempty"`
			Reputation      float64 `json:"reputation,omitempty"`
			SandboxVerdicts struct {
				C2AE struct {
					Category              string   `json:"category,omitempty"`
					MalwareClassification []string `json:"malware_classification,omitempty"`
					SandboxName           string   `json:"sandbox_name,omitempty"`
				}
			} `json:"sandbox_verdicts,omitempty"`
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
					} `json:"values,omitempty"`
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
			} `json:"sigma_analysis_stats,omitempty"`
			SigmaAnalysisSummary struct {
				Sigma_Integrated_Rule_Set__GitHub_ struct {
					Critical float64 `json:"critical,omitempty"`
					High     float64 `json:"high,omitempty"`
					Low      float64 `json:"low,omitempty"`
					Medium   float64 `json:"medium,omitempty"`
				} `json:"Sigma Integrated Rule Set (GitHub),omitempty"`
			} `json:"sigma_analysis_summary,omitempty"`
			Size           float64  `json:"size,omitempty"`
			Ssdeep         string   `json:"ssdeep,omitempty"`
			Tags           []string `json:"tags,omitempty"`
			TimesSubmitted float64  `json:"times_submitted,omitempty"`
			Tlsh           string   `json:"tlsh,omitempty"`
			TotalVotes     struct {
				Harmless  float64 `json:"harmless,omitempty"`
				Malicious float64 `json:"malicious,omitempty"`
			} `json:"total_votes,omitempty"`
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
		} `json:"attributes,omitempty"`
		ID    string `json:"id,omitempty"`
		Links struct {
			Self string `json:"self,omitempty"`
		} `json:"links,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"data,omitempty"`
}
