package commands

import (
	"bytes"
	"runtime"
	"text/template"
	"time"

	"github.com/griffin702/service/tools"
)

var (
	Version   = "v1.0.0"
	BuildTime = tools.Tools.TimeFormat(time.Now(), "Y-m-d H:i:s")
)

// VersionOptions include version
type VersionOptions struct {
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = `Version: {{.Version}} | Go version: {{.GoVersion}} | BuildTime: {{.BuildTime}} | OS/Arch: {{.Os}}/{{.Arch}}`

func GetVersion() string {
	var doc bytes.Buffer
	vo := VersionOptions{
		Version:   Version,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
	tmpl, _ := template.New("version").Parse(versionTemplate)
	_ = tmpl.Execute(&doc, vo)
	return doc.String()
}
