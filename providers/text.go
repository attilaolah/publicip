package providers

import (
	"io"
	"io/ioutil"
	"net"
	"strings"
)

func getText(url string) (net.IP, error) {
	res, err := get(url, "text/html")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return parseText(res.Body)
}

func parseText(src io.Reader) (net.IP, error) {
	body, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return parseIP(string(body))
}

func parseIP(s string) (net.IP, error) {
	if ip := net.ParseIP(strings.TrimSpace(s)); ip != nil {
		return ip, nil
	}

	return nil, NotFound
}
