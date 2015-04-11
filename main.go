package main

import (
	"github.com/codegangsta/cli"
	"github.com/mgutz/logxi/v1"
	"os"
	"path/filepath"
)

var logger log.Logger

var (
	BasePath   string
	SrcPath    string
	TargetPath string
)

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
	if os.Getenv("LOGXI") == "" {
		log.ProcessLogxiEnv("*=INF")
	}
	logger = log.New("linen")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Error(err.Error())
	}
	BasePath = dir
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

	if os.Getenv("LOGXI") == "" {
		os.Setenv("LOGXI", "*=INF")
	}

	app.Run(os.Args)
}

func BuildAction(c *cli.Context) {
	SrcPath = c.String("src")
	TargetPath = c.String("target")

	Build(SrcPath, TargetPath)

	logger.Info("DONE")
}

func PreviewAction(c *cli.Context) {
	var err error

	SrcPath, err = filepath.Abs(c.String("src"))
	if err != nil {
		logger.Error("srcPath", "path", SrcPath, "err", err)
		os.Exit(1)
	}
	TargetPath, err = filepath.Abs(c.String("target"))
	if err != nil {
		logger.Error("targetPath", "path", TargetPath, "err", err)
		os.Exit(1)
	}

	build := c.Bool("build")
	addr := c.String("addr")

	previewServe(addr, build)

}
