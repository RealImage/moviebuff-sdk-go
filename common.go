package moviebuff

import (
	"net/http"
)

const apiKey = "X-Api-Key"

func prepareRequest(hostURL, staticToken, path string) (*http.Request, error) {

	r, err := http.NewRequest(http.MethodGet, hostURL+path, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add(apiKey, staticToken)
	return r, nil
}
