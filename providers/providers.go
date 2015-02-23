package providers

import "net"

func JSONIP() (net.IP, error) {
	return getJSON("http://jsonip.org")
}

func APIFy() (net.IP, error) {
	return getText("https://api.ipify.org")
}
