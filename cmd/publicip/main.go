package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/attilaolah/publicip"
)

func main() {
	n := flag.Bool("n", false, "no newline")
	flag.Parse()

	ip, err := publicip.IP()
	if err != nil {
		fmt.Println("publicip: %v", err)
		os.Exit(1)
	}

	if ip == nil {
		panic("ip was nil")
	}

	if *n {
		fmt.Print(ip)
	} else {
		fmt.Println(ip)
	}
}
