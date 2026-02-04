package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/p9works/p9/internal/ports"
	"net"
	"syscall"
	"time"
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
	timeoutFlag := flag.Duration("t", 3*time.Second, "Override default timeout (e.g. -t 5s, -t 60s)")
	//localFlag := flag.Bool("l", false, "List local open ports")
	//domainFlag := flag.String("d", "", "Domain/IP lookup")
	flag.Parse()

	switch {
	case *remoteFlag != "":
		handleRemote(*remoteFlag, *timeoutFlag)
	default:
		printUsage()
	}
}

func handleRemote(address string, timeout time.Duration) {
	isOpen, err := ports.CheckPortTCP(address, timeout)
	if err != nil {
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			var netErr net.Error
			if errors.As(opErr.Err, &netErr) && netErr.Timeout() {
				fmt.Printf("⏱️  \033[33mConnection to %s TIMED OUT\033[0m\n", address)
				return
			}
			if errors.Is(opErr.Err, syscall.ECONNREFUSED) {
				fmt.Printf("🔴 \033[31mPort %s is CLOSED (connection refused)\033[0m\n", address)
				return
			}
			var dnsErr *net.DNSError
			if errors.As(opErr.Err, &dnsErr) {
				fmt.Printf("🔍 \033[33mDNS lookup failed: %v\033[0m\n", dnsErr)
				return
			}
		}
		var addrErr *net.AddrError
		if errors.As(err, &addrErr) {
			fmt.Printf("⚠️  \033[33mInvalid address format: %v\033[0m\n", addrErr)
			return
		}
		fmt.Printf("❌ \033[31mCannot connect to %s: %v\033[0m\n", address, err)
		return
	}
	if isOpen {
		fmt.Printf("✅ \033[32mPort %s is OPEN\033[0m\n", address)
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
