package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "PRODUCER"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start consumer [--type {simple|async|batch} --topic]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "type",
					Value: "simple",
					Usage: "Choose the consume type simple|orderly|broadcast",
				},
				&cli.StringFlag{
					Name:  "topic",
					Value: "simple",
					Usage: "insert a topic name",
				},
			},
			Action: func(c *cli.Context) (err error) {
				switch c.String("type") {
				case "simple":
					err = Simple(c.String("topic"))
				case "async":
					err = Async(c.String("topic"))
				case "tag":
					err = Tag(c.String("topic"))

				}
				if err != nil {
					return
				}
				return nil
			},
		},
	}
	cliApp.Run(os.Args)
}
