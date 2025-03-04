package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/chaosannals/goctl-ts/action"
	"github.com/urfave/cli/v2"
)

var (
	commands = []*cli.Command{
		{
			Name:   "ts",
			Usage:  "generates http client for ts",
			Action: action.Ts,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "caller",
					Usage:    "the caller of ts",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "url",
					Usage:    "custom url prefix.",
					Required: false,
				},
				&cli.BoolFlag{
					Name:  "unwrap",
					Usage: "unwrap api client",
					Value: false,
				},
				&cli.BoolFlag{
					Name:  "body",
					Usage: "add custom body param.",
					Value: false,
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate http client code for ts."
	app.Version = fmt.Sprintf("%s %s/%s", action.Version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-ts: %+v\n", err)
	}
}
