package main

import (
	"fmt"
	"gotwt/twitter"
	"gotwt/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	tclient := twitter.GetClient(config.ApiKey, config.ApiKeySecret)

	retweeters, errs1 := twitter.GetRetweets(tclient, "1407778348441296899")
	if errs1 != nil {
		panic(errs1)
	}
	fmt.Println(retweeters)

	liking_users, errs2 := twitter.GetLikes(tclient, "1407778348441296899")
	if errs2 != nil {
		panic(errs2)
	}
	fmt.Println(liking_users)
}
