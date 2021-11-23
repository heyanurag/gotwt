package main

import (
	"gotwt/util"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	// tclient := twitter.GetClient(config.ApiKey, config.ApiKeySecret)

	// retweeters, errs1 := twitter.GetRetweets(tclient, "1407778348441296899")
	// if errs1 != nil {
	// 	panic(errs1)
	// }
	// fmt.Println(retweeters)

	// liking_users, errs2 := twitter.GetLikes(tclient, "1407778348441296899")
	// if errs2 != nil {
	// 	panic(errs2)
	// }
	// fmt.Println(liking_users)

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

			// fmt.Printf("Hello %q, %t, %t, %d", c.Args().Get(0), c.Bool("retweets"), c.Bool("likes"), c.Int("winners"))

			util.GetWinners(c.Args().Get(0), c.Int("winners"), c.Bool("retweets"), c.Bool("likes"))

			return nil
		},
	}

	err3 := app.Run(os.Args)
	if err3 != nil {
		log.Fatal(err3)
	}
}
