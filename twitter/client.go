package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

func GetClient(apiKey string, apiSecretKey string) *http.Client {
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token?grant_type=client_credentials", strings.NewReader(""))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(apiKey, apiSecretKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoding;charset=UTF-8")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var token oauth2.Token
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&token)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	conf := &oauth2.Config{}
	tclient := conf.Client(context.Background(), &token)

	return tclient
}
