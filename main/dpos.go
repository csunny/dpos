package main

import (
	"os"
	"github.com/outbrain/golib/log"
	"github.com/urfave/cli"
	"github.com/csunny/dpos"
	"github.com/csunny/dpos/tools"
)

const (
	// Version 定义
	Version = "v0.1.0"
)

var mainFlags = []cli.Flag{
	cli.StringFlag{
		Name: "log_level",
		Value: "INFO",
		Usage: "set the log level to INFO",
	},
}

var mainCommands = []cli.Command{
	tools.NodeVote,
	dpos.NewNode,
	
}

func main()  {
	app := cli.NewApp()
	app.Name = "dpos"
	app.Version = Version
	app.Author = "Magic"
	app.Flags = mainFlags
	app.Usage = "a simple DPoS algorithm implement"
	app.Commands = mainCommands
	app.Before = func(context *cli.Context) error {
		if context.GlobalBool("version") {
			log.Info(Version)
			os.Exit(0)
		}
		logLevel,err := log.LogLevelFromString(context.GlobalString("log_level"))
		if err != nil {
			log.Errorf("unknown log level:%s",err.Error())
		} else {
			log.SetLevel(logLevel)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Errorf(err.Error())		
	}
}