package main

import (
	"fmt"
	"os"

	"github.com/attilaolah/publicip"
)

func main() {
	ip, err := publicip.IP()
	if err != nil {
		fmt.Errorf("publicip: %v", err)
		os.Exit(1)
	}

	if ip == nil {
		panic("ip was nil")
	}

	fmt.Print(ip)
}
