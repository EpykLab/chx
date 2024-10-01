package pretty

import "errors"

// Handles printing the content to standard out
// in pretty format
//
//	dtype: data type, such as IP, domain, or hash,
//	source: source of data such as alienvault or crowdsec,
//	content: json content that will be parsed and pretty printed
func PrintContentPretty(dtype string, source string, content string) error {
	switch dtype {
	case "domain":
		switch source {
		case "alienvault":
			if err := parseAVIPContent(); err != nil {
				return err
			}
		}
	case "hash":
		switch source {
		case "vt":
			if err := parseAVTHashContent(); err != nil {
				return err
			}
		case "alienvault":
			if err := parseAVTHashContent(); err != nil {
				return err
			}
		}
	case "ip":
		switch source {
		case "alienvault":
			if err := parseAVIPContent(); err != nil {
				return err
			}
		case "aipdb":
			if err := parseAIPDBIPContent(); err != nil {
				return err
			}
		case "crowdsec":
			if err := parseCSIPContent(); err != nil {
				return err
			}
		}
	default:
		return errors.New("option unknown")
	}

	return nil
}
