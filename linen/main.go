package main

import (
    "os"
//    "path/filepath"
    "github.com/codegangsta/cli"
    "github.com/anarcher/linen"
)

func main() {
    app := cli.NewApp()
    app.Name = "linen"
    app.Usage = "A Simple Static Site Generator in Go"
    app.Version = "0.0.1"

    app.Flags = []cli.Flag {
        cli.StringFlag{"config, c", "./linen.yaml", "config file (default is current path/linen.yaml|json|toml)"},
        cli.StringFlag{"root, r", ".", "root path (default is current path"},
    }

    app.Commands = []cli.Command{
        {
            Name: "build",
            ShortName: "b",
            Usage: "Build source files for render",
            Action: func(c *cli.Context) {
                config,err := linen.NewConfig(c.GlobalString("config"))
                if err != nil {
                    println("Failed to read config file:",err)
                    return
                }
                println(config.ConfigFile)
                //println(c.GlobalString("config"))
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
