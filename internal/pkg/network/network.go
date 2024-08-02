// Package network provides network utilities.
package network

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

// GetLocalIP returns the non loopback local IP of the host.
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// Check the address type and, if it is not a loopback, compare it with ip.
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("failed to get local ip address")
}

// IsLocalIP reports whether ip is the IP of the host.
func IsLocalIP(ip string) bool {
	// Check whether ip is one of known local addresses.
	if ip == "127.0.0.1" || ip == "localhost" || ip == "::1" {
		return true
	}

	localIP, err := GetLocalIP()
	if err != nil {
		log.Err(err).Msg("failed to get local ip")
		return false
	}

	return ip == localIP
}
