package fetcher

import (
	"net/http"
	"testing"
)

const url = "http://album.zhenai.com/u/1764131916"

func TestFetch(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error(resp.StatusCode)
	}
}
