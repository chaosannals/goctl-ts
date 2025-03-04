package generate

import "github.com/chaosannals/goctl-ts/template"

func GenRequests(dir string, caller string) error {
	data := template.RequestTemplateData{
		Caller: caller,
	}
	return template.GenTsFile(dir, "gocliRequest", template.Requests, data)
}
