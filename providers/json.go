package providers

import (
	"encoding/json"
	"io"
	"net"
)

func getJSON(url string) (net.IP, error) {
	res, err := get(url, "application/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return parseJSON(res.Body)
}

func parseJSON(src io.Reader) (net.IP, error) {
	res := struct {
		IP string `json:"ip"`
	}{}
	if err := json.NewDecoder(src).Decode(&res); err != nil {
		return nil, err
	}

	return parseIP(res.IP)
}
