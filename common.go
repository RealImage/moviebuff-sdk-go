package moviebuff

import (
	"net/http"
	"os"
)

var (
	hostUrl = "https://api.moviebuff.com/api/v2"
)

func init() {
	if os.Getenv("MOVIEBUFF_URL") != "" {
		hostUrl = os.Getenv("MOVIEBUFF_URL")
	}
}

type logger interface {
	Println(v ...interface{})
}

// Dummy discard, satisfies io.Writer without importing io or os.
type devNull struct{}

func (devNull) Write(p []byte) (int, error) {
	return len(p), nil
}

func prepareRequest(hostURL, staticToken, path string) (*http.Request, error) {
	if hostURL == "" {
		hostURL = hostUrl
	}

	r, err := http.NewRequest(http.MethodGet, hostURL+path, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("X-Api-Key", staticToken)
	return r, nil
}
