package twitter

import (
	"io"
	"net/http"
	"os"
)

func GetRetweets(tclient *http.Client, tweetID string) (err error) {
	url := "https://api.twitter.com/2/tweets/" + tweetID + "/retweeted_by"
	retweets, err := tclient.Get(url)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, retweets.Body)

	return nil
}
