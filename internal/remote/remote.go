package remote

import (
	"net"
	"time"
)

func CheckPortTCP(address string) (bool, error) {
	result, err := net.DialTimeout("tcp", address, 3*time.Second)

	if err != nil {
		return false, err
	}
	defer result.Close()
	return true, nil
}
