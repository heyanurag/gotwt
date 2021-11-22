package twitter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetRetweets(tclient *http.Client, tweetID string) (res Response, err error) {
	url := "https://api.twitter.com/2/tweets/" + tweetID + "/retweeted_by"
	retweets, err := tclient.Get(url)
	if err != nil {
		return
	}
	defer retweets.Body.Close()

	byteValue, _ := ioutil.ReadAll(retweets.Body)
	var response Response
	json.Unmarshal(byteValue, &response)

	return response, nil
}
