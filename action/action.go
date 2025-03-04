package action

import (
	"github.com/chaosannals/goctl-ts/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

const Version = "20250304"

func Ts(ctx *cli.Context) error {
	caller := ctx.String("caller")
	unwrap := ctx.Bool("unwrap")
	body := ctx.Bool("body")
	url := ctx.String("url")

	plugin, err := plugin.NewPlugin()
	if err != nil {
		return err
	}

	return generate.TsCommand(plugin, caller, unwrap, body, url, Version)
}
