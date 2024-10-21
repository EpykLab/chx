package pretty

import (
	"errors"

	"github.com/EpykLab/chx/cmd/utils/pretty/data"
)

type markdownPrinter interface {
	printer(d any) error
}

// Handles printing the content to standard out
// in pretty format
//
//	dtype: data type, such as IP, domain, or hash,
//	source: source of data such as alienvault or crowdsec,
//	content: json content that will be parsed and pretty printed
func PrintContentPretty(dtype data.Dtype, source data.Source, content interface{}) error {
	switch dtype {
	case data.Domain:
		switch source {
		case data.AlienVault:
			d := content.(*AlienVaultDomain)
			if err := d.printer(d); err != nil {
				return err
			}
		}
	case data.Hash:
		switch source {
		case data.VirusTotal:
			d := content.(*VirusTotalHash)
			if err := d.printer(d); err != nil {
				return err
			}
		case data.AlienVault:
			d := content.(*AlienVaultHash)
			if err := d.printer(d); err != nil {
				return err
			}
		}
	case data.IP:
		switch source {
		case data.AlienVault:
			d := content.(*AlientVaultIP)
			if err := d.printer(d); err != nil {
				return err
			}
		case data.IpAbuseDB:
			d := content.(*AipdbIP)
			if err := d.printer(d); err != nil {
				return err
			}
		case data.CrowdSec:
			d := content.(*CrowdSecIP)
			if err := d.printer(d); err != nil {
				return err
			}
		}
	default:
		return errors.New("option unknown")
	}

	return nil
}
