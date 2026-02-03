package main

import (
	"flag"
	"fmt"
	"github.com/p9works/p9/internal/remote"
)

func main() {
	// TODO: Define flags here
	/*
		1. Open ports and listening host - "p9 -l",
		2. Check remote port  - check remote tcp and udp ports "p9 -r 10.10.1.1:8080"
		3. domain whois and ip lookup - "p9 -d example.com"
	*/

	// TODO: Parse flags

	// TODO: Check which flag is set and route to the right function

	// For now, just print which operation was requested
	remoteFlag := flag.String("r", "", "Check remote port (host:port)")
	//localFlag := flag.Bool("l", false, "List local open ports")
	//domainFlag := flag.String("d", "", "Domain/IP lookup")
	flag.Parse()

	switch {
	case *remoteFlag != "":
		handleRemote(*remoteFlag)
	default:
		printUsage()
	}

	//if *remoteFlag != "" {
	//	isOpen, err := remote.CheckPortTCP(*remoteFlag)
	//	if err != nil {
	//		fmt.Printf("Cannot connect to %s, %v\n", *remoteFlag, err)
	//		return
	//	}
	//	if isOpen {
	//		fmt.Printf("✓ Port %s is OPEN\n", *remoteFlag)
	//	} else {
	//		fmt.Printf("✗ Port %s is CLOSED\n", *remoteFlag)
	//	}
	//}
}

func handleRemote(address string) {
	isOpen, err := remote.CheckPortTCP(address)
	if err != nil {
		fmt.Printf("Cannot connect to %s, %v\n", address, err)
		return
	}
	if isOpen {
		fmt.Printf("✓ Port %s is OPEN\n", address)
	} else {
		fmt.Printf("✗ Port %s is CLOSED\n", address)
	}
}

func printUsage() {
	fmt.Println("Debugging tool")
	fmt.Println("\nUsage:")
	fmt.Println("  p9 -r <host:port>  Check if remote port is open")
	fmt.Println("  p9 -l              List local open ports")
	fmt.Println("  p9 -d <domain>     Lookup domain/IP information")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}
