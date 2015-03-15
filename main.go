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
	}

	app.Run(os.Args)
}

func BuildAction(c *cli.Context) {
	srcPath := c.String("src")
	targetPath := c.String("target")

	var err error
	var files Files

	files, err = ReadFiles(srcPath)
	if err != nil {
		logger.Error("Read", "err", err)
	}

	err = TransformFiles(files)
	if err != nil {
		logger.Error("Transform", "err", err)
	}

	err = WriteFiles(files, targetPath)
	if err != nil {
		logger.Error("Write", "err", err)
	}

	logger.Info("DONE")
}
