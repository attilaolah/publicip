package publicip

import (
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/attilaolah/publicip/providers"
)

var (
	// How long to cache thet IP?
	Timeout = time.Minute

	cache struct {
		net.IP
		time.Time
		sync.Mutex
	}
)

// IP returns the public IP address.
//
// Calling multiple times will result in a cached value being returned.
// To disable caching, set publicip.Timeout to zero.
func IP() (net.IP, error) {
	cache.Lock()
	defer cache.Unlock()

	if time.Since(cache.Time) < Timeout {
		return cache.IP, nil
	}

	errs := []string{}
	for provider, fn := range providers.All {
		res, err := fn()
		if err == nil {
			cache.IP = res
			cache.Time = time.Now()
			return res, nil
		}
		errs = append(errs, provider+": "+err.Error())
	}
	return nil, errors.New(strings.Join(errs, "; "))
}

// Refresh returns the public IP, regardless of the state of the cache.
func Refresh() (net.IP, error) {
	cache.Lock()
	cache.Time = time.Time{}
	cache.Unlock()

	return IP()
}
