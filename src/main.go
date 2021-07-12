//目前是同步链路版
package main

import (
	"fmt"
	"log"
	"os"
	"yond/file"
	"yond/funcs"
	"yond/util"

	"github.com/urfave/cli/v2"
)

func initSelf(rootPath string) {
	util.RootPath = rootPath
	content := file.ReadFile(rootPath + `\dev\js.mod`)
	matches := util.Re_Js_Mod.FindAllString(content, -1)
	for _, value := range matches {
		util.SetModuleUrl[util.Re_between_quotation.FindString(value)] = true
	}
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init a yond project",
			Action: func(c *cli.Context) error {
				rootPath, _ := os.Getwd()
				initSelf(rootPath)
				funcs.Init()
				return nil
			},
		},
		{
			Name:    "dev",
			Aliases: []string{"d"},
			Usage:   "transform http import and create dev directory",
			Action: func(c *cli.Context) error {
				rootPath, _ := os.Getwd()
				initSelf(rootPath)
				funcs.BuildDev(rootPath)
				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build and create dist directory",
			Action: func(c *cli.Context) error {
				rootPath, _ := os.Getwd()
				initSelf(rootPath)
				fmt.Println("build with esbuild", c.Args().First())
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
