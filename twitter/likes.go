package twitter

import (
	"io"
	"net/http"
	"os"
)

func GetLikes(tclient *http.Client, tweetID string) (err error) {
	url := "https://api.twitter.com/2/tweets/" + tweetID + "/liking_users"
	likes, err := tclient.Get(url)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, likes.Body)

	return nil
}
