package moviebuff

import (
	"net/http"
)

func prepareRequest(hostURL, staticToken, path string) (*http.Request, error) {

	r, err := http.NewRequest(http.MethodGet, hostURL+path, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("X-Api-Key", staticToken)
	return r, nil
}
