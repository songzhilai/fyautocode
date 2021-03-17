package main

import (
	"github.com/songzhilai/fyautocode/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//加载静态文件至二进制
//go-bindata -o=./asset/asset.go -pkg=asset template/...

func main() {
	app := cli.NewApp()
	app.Name = "FYAPI"
	app.Usage = "API脚手架"
	app.Version = commands.GetVersion()
	app.Authors = []*cli.Author{{
		Name:  "FY",
		Email: "781502032@qq.com",
	}}
	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help",
		Usage: "查看帮助",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version,v",
		Aliases: []string{"v"},
		Usage:   "GoCMD Version",
	}
	// app.StringFlag == &cli.StringFlag{
	// 	Name:  "type,t",
	// 	Usage: "设置服务类型",
	// }

	app.Commands = commands.InitCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}
