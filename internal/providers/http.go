package providers

import "net/http"

func get(url, accept string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", accept)

	return http.DefaultClient.Do(req)
}
