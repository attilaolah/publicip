package publicip

import (
	"errors"
	"net"
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

	ip, err := fastIP()
	if err != nil {
		return nil, err
	}

	cache.IP = ip
	cache.Time = time.Now()
	return ip, nil
}

// Refresh returns the public IP, regardless of the state of the cache.
func Refresh() (net.IP, error) {
	cache.Lock()
	cache.Time = time.Time{}
	cache.Unlock()

	return IP()
}

// fastIP fires up a goroutine for each provider, returning an IP as soon as one is found.
func fastIP() (net.IP, error) {
	ips := make(chan net.IP, 1)
	errs := make(chan error)
	wg := sync.WaitGroup{}

	wg.Add(len(providers.All))
	for provider, fn := range providers.All {
		go func(provider string, fn func() (net.IP, error)) {
			if ip, err := fn(); err == nil {
				ips <- ip
			} else {
				errs <- err
			}
			wg.Done()
		}(provider, fn)
	}

	go func() {
		wg.Wait()
		close(ips)
		close(errs)
	}()

	for ip := range ips {
		return ip, nil
	}

	for err := range errs {
		return nil, err
	}

	return nil, errors.New("no providers")
}
