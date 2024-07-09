package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type ChapiteauArguments struct {
	Path       string
	ServiceUrl string
	Auth       string
	Project    string
	BuildId    string
	BuildName  string
	ReportURL  string
}

func main() {
	app := &cli.App{
		Name:  "chapiteau-cli",
		Usage: "Upload report/data to the chapiteau service",
		Commands: []*cli.Command{
			{
				Name:  "upload",
				Usage: "Upload a report or just index.html file",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "path", Usage: "Path to the folder or file to upload", Required: true},
					&cli.StringFlag{Name: "url", Usage: "Chapiteau service url", Required: true},
					&cli.StringFlag{Name: "auth", Usage: "Authentication token", Required: true},
					&cli.StringFlag{Name: "project", Usage: "Project ID", Required: true},
					&cli.StringFlag{Name: "build-id", Usage: "CI Build ID"},
					&cli.StringFlag{Name: "build-name", Usage: "CI Build Name"},
					&cli.StringFlag{Name: "report-url", Usage: "Playwright Report URL hosted elsewhere"},
				},
				Action: func(c *cli.Context) error {
					args := ChapiteauArguments{
						Path:       c.String("path"),
						ServiceUrl: c.String("url"),
						Auth:       c.String("auth"),
						Project:    c.String("project"),
						BuildId:    c.String("build-id"),
						BuildName:  c.String("build-name"),
						ReportURL:  c.String("report-url"),
					}

					isDir, err := isDirectory(args.Path)
					if err != nil {
						return fmt.Errorf("failed to check if path is a directory: %s", err)
					}

					if isDir {
						return uploadReport(args)
					} else {
						return uploadIndexHtml(args)
					}
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func isDirectory(path string) (isDir bool, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}
