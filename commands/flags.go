package commands

import "github.com/urfave/cli/v2"

type Flags struct {
	Name      string
	Path      string
	ModPrefix string
	Type      string
}

var f *Flags

func (f *Flags) ToNewAction() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "d",
			Value:       "",
			Usage:       "指定项目所在目录",
			Destination: &f.Path,
		},
		&cli.StringFlag{
			Name:        "t",
			Value:       "apisrv",
			Usage:       "服务类型",
			Destination: &f.Type,
		},
	}
}
