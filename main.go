package main

import (
	"gotwt/util"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "gotwt",
		Usage: "Find twitter giveaway winners!!",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "winners",
				Aliases: []string{"w"},
				Usage:   "Total number of winners",
				Value:   1,
			},
			&cli.BoolFlag{
				Name:    "retweets",
				Aliases: []string{"r"},
				Usage:   "Find winners from retweets",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "likes",
				Aliases: []string{"l"},
				Usage:   "Find winners from likes",
				Value:   false,
			},
		},
		Action: func(c *cli.Context) error {

			util.GetWinners(c.Args().Get(0), c.Int("winners"), c.Bool("retweets"), c.Bool("likes"))

			return nil
		},
	}

	err3 := app.Run(os.Args)
	if err3 != nil {
		log.Fatal(err3)
	}
}
