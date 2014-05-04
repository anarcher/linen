package main

import (
    "os"
//    "path/filepath"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "linen"
    app.Usage = "A Simple Static Site Generator in Go"
    app.Version = "0.0.1"

    currentDir := "."

    app.Flags = []cli.Flag {
        cli.StringFlag{"config, c", currentDir+"/linen.yaml", "config file (default is current path/linen.yaml|json|toml)"},
        cli.StringFlag{"root, r", currentDir, "root path (default is current path"},
    }

    app.Commands = []cli.Command{
        {
            Name: "build",
            ShortName: "b",
            Usage: "Build source files for render",
            Action: func(c *cli.Context) {
                println(c.GlobalString("config"))
            },
            Flags: []cli.Flag {
                cli.StringFlag{"text","plain","text format"},
            },
        },
        {
            Name: "server",
            ShortName: "s",
            Usage: "Runs it's own a webserver to render the files",
            Action: func(c *cli.Context) {
                println("Serve...")
            },
            Flags:[]cli.Flag {
                cli.BoolTFlag{ "watch, w","watch files for changes and recreate as needed"},
            },
        },
    }

    app.Run(os.Args)
}
