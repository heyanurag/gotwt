package util

import (
	"fmt"
	"gotwt/twitter"
	"log"
	"math/rand"
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
	// fmt.Println(retweeters)

	liking_users, errs2 := twitter.GetLikes(tclient, tweetid)
	if errs2 != nil {
		panic(errs2)
	}
	// fmt.Println(liking_users)

	var winners []twitter.User

	if fromLikes && fromRetweets {
		pool := append(retweeters.Data, liking_users.Data...)
		winners = pickWinners(pool, noOfWinners)
	} else if fromLikes {
		winners = pickWinners(liking_users.Data, noOfWinners)
	} else if fromRetweets {
		winners = pickWinners(retweeters.Data, noOfWinners)
	}

	fmt.Println(winners)

}

func pickWinners(pool []twitter.User, noOfWinners int) []twitter.User {
	fmt.Println(len(pool))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(pool))

	winners := perm[:noOfWinners]

	ret := make([]twitter.User, 0, noOfWinners)
	for _, idx := range winners {
		ret = append(ret, pool[idx])
	}
	return ret
}
