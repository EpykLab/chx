package pretty

import (
	"embed"
	"os"
	"text/template"

	"github.com/charmbracelet/log"
)

//go:embed static/layouts/*

var StaticFiles embed.FS

type File string

const (
	AIP      File = "aip"
	AV       File = "av"
	AVIP     File = "av-ip"
	AVHASH   File = "av-hash"
	AVDOMAIN File = "av-domain"
	CS       File = "cs"
	VT       File = "vt"
)

type RenderEngine interface {
	New() (*Output, error)
	Render(f File) error
}

type Output struct {
	Templ template.Template
}

func New() (*Output, error) {
	tmpl, err := template.ParseFS(StaticFiles, "**/**/*.gomd")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Output{
		Templ: *tmpl,
	}, nil
}

func (o *Output) Render(f File, d any) error {
	x := o.Templ

	err := x.ExecuteTemplate(os.Stdout, string(f), d)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
