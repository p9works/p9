package ports

import (
	"net"
	"time"
)

func CheckPortTCP(address string, timeout time.Duration) (bool, error) {
	result, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		return false, err
	}
	defer result.Close()
	return true, nil
}
