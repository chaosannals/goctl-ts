package generate

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func TsCommand(
	p *plugin.Plugin,
	caller string,
	unwrap bool,
	body bool,
	url string,
	version string,
) error {
	api, err := parser.Parse(p.ApiFilePath)
	if err != nil {
		fmt.Println(color.Red.Render("Failed"))
		return err
	}

	if err := api.Validate(); err != nil {
		return err
	}

	if len(caller) == 0 {
		caller = "webapi"
	}

	api.Service = api.Service.JoinPrefix()
	logx.Must(pathx.MkdirIfNotExist(p.Dir))
	logx.Must(GenRequests(p.Dir, caller))
	logx.Must(GenHandler(p.Dir, caller, api, unwrap, body, url))
	logx.Must(GenComponents(p.Dir, api, version))

	fmt.Println(color.Green.Render("Done."))
	return nil
}
