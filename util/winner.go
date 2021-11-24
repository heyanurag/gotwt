package util

import (
	"fmt"
	"gotwt/twitter"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetWinners(tweetid string, noOfWinners int, fromRetweets bool, fromLikes bool) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	tclient := twitter.GetClient(config.ApiKey, config.ApiKeySecret)

	retweeters, errs1 := twitter.GetRetweets(tclient, tweetid)
	if errs1 != nil {
		panic(errs1)
	}

	liking_users, errs2 := twitter.GetLikes(tclient, tweetid)
	if errs2 != nil {
		panic(errs2)
	}

	var winners []twitter.User

	if fromLikes && fromRetweets {
		pool := append(retweeters.Data, liking_users.Data...)
		winners = pickWinners(pool, noOfWinners)
	} else if fromLikes {
		winners = pickWinners(liking_users.Data, noOfWinners)
	} else if fromRetweets {
		winners = pickWinners(retweeters.Data, noOfWinners)
	}

	for i := 0; i < len(winners); i++ {
		fmt.Printf("%d. %s  (@%s)\n", i+1, winners[i].Name, winners[i].Username)
	}

	fmt.Println("\nCongratulations to all the winner(s)!🎉")

}

func pickWinners(pool []twitter.User, noOfWinners int) []twitter.User {

	if noOfWinners > len(pool) {
		fmt.Printf("\nOops! Cannot pick %d out of %d participants😕\n\n", noOfWinners, len(pool))
		os.Exit(1)
	}

	fmt.Printf("\nPicking winners from a pool of %d participants!🌈\n\n", len(pool))

	saveToExcel(pool)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(pool))

	winners := perm[:noOfWinners]

	ret := make([]twitter.User, 0, noOfWinners)
	for _, idx := range winners {
		ret = append(ret, pool[idx])
	}

	return ret
}
