package main

import (
	"github.com/codegangsta/cli"
	"github.com/mgutz/logxi/v1"
	"os"
)

var logger log.Logger

var (
	flSrcPath = cli.StringFlag{
		Name:  "src,s",
		Value: ".",
		Usage: "",
	}
	flTargetPath = cli.StringFlag{
		Name:  "target,t",
		Value: "./target",
		Usage: "",
	}
	flBuild = cli.BoolFlag{
		Name:  "build,b",
		Usage: "",
	}
	flAddr = cli.StringFlag{
		Name:  "addr,a",
		Value: ":8080",
		Usage: "preview address",
	}
)

func init() {
	logger = log.New("linen")
	logger.SetLevel(log.LevelInfo)
}

func main() {

	app := cli.NewApp()
	app.Name = "linen"
	app.Usage = "is simple static page(s) generator"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:      "build",
			ShortName: "b",
			Usage:     "build source pages",
			Flags:     []cli.Flag{flSrcPath, flTargetPath},
			Action:    BuildAction,
		},
		{
			Name:      "preview",
			ShortName: "p",
			Usage:     "preview",
			Flags:     []cli.Flag{flSrcPath, flTargetPath, flBuild, flAddr},
			Action:    PreviewAction,
		},
	}

	app.Run(os.Args)
}

func BuildAction(c *cli.Context) {
	srcPath := c.String("src")
	targetPath := c.String("target")

	Build(srcPath, targetPath)

	logger.Info("DONE")
}

func PreviewAction(c *cli.Context) {
	srcPath := c.String("src")
	targetPath := c.String("target")
	build := c.Bool("build")
	addr := c.String("addr")

	previewServe(addr, srcPath, targetPath, build)

}
