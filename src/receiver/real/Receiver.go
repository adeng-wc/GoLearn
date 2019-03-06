package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Receiver struct {
	UserAgent string
	TimeOut   time.Duration
}

//如果 Receiver 很大，使用值传递就会拷贝,
func (r *Receiver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
