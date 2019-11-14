package moviebuff

import (
	"context"
	"net/http"
)

const apiKey = "X-Api-Key"

func prepareRequest(ctx context.Context, hostURL, staticToken, path string) (*http.Request, error) {

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, hostURL+path, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add(apiKey, staticToken)
	return r, nil
}

func addQueryParams(r *http.Request, queryParams map[string]string) {
	q := r.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
}
