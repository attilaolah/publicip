package providers

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
)

var NotFound = errors.New("public address not found")

func JSONIP() (net.IP, error) {
	res, err := http.Get("http://jsonip.org")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	jres := struct {
		IP string `json:"ip"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&jres); err != nil {
		return nil, err
	}

	ip := net.ParseIP(jres.IP)
	if ip == nil {
		return nil, NotFound
	}
	return ip, nil
}
