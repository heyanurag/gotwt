package twitter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetLikes(tclient *http.Client, tweetID string) (res Response, err error) {
	url := "https://api.twitter.com/2/tweets/" + tweetID + "/liking_users"
	likes, err := tclient.Get(url)
	if err != nil {
		return
	}
	defer likes.Body.Close()

	byteValue, _ := ioutil.ReadAll(likes.Body)
	var response Response
	json.Unmarshal(byteValue, &response)

	return response, nil
}
