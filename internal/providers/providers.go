package providers

import "net"

var (
	// All contains a list of known providers.
	All = map[string]provider{}

	// HTTP contains a list of known providers that accept HTTP connections.
	HTTP = map[string]provider{}

	// HTTPS contains a list of known providers that accept HTTPS connections.
	HTTPS = map[string]provider{}
)

type provider func() (net.IP, error)

func textProvider(url string) provider {
	return func() (net.IP, error) { return getText(url) }
}

func jsonProvider(url string) provider {
	return func() (net.IP, error) { return getJSON(url) }
}

func init() {
	// HTTPS providers:
	for _, url := range []string{
		"https://api.ipify.org",
		"https://icanhazip.com",
		"https://ip.appspot.com",
		"https://ipinfo.io/ip",
	} {
		HTTPS[url] = textProvider(url)
	}

	// HTTP providers:
	for _, url := range []string{
		"http://api.ipify.org",
		"http://curlmyip.com",
		"http://icanhazip.com",
		"http://ip.appspot.com",
		"http://ipinfo.io/ip",
	} {
		HTTP[url] = textProvider(url)
	}
	for _, url := range []string{
		"http://ifconfig.me/ip",
		"http://jsonip.org",
	} {
		HTTP[url] = jsonProvider(url)
	}

	for key, val := range HTTP {
		All[key] = val
	}
	for key, val := range HTTPS {
		All[key] = val
	}
}
