package netx

import (
	"net"
	"net/http"
	"strings"
)

// IsInternalIP verify an ip is intranet or not.
func IsInternalIP(IP net.IP) bool {
	if IP.IsLoopback() {
		return true
	}
	if ip4 := IP.To4(); ip4 != nil {
		return ip4[0] == 10 ||
			(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
			(ip4[0] == 169 && ip4[1] == 254) ||
			(ip4[0] == 192 && ip4[1] == 168)
	}
	return false
}

// GetRequestPublicIp return the requested public ip.
func GetRequestPublicIp(req *http.Request) string {
	var ip string
	for _, ip = range strings.Split(req.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !IsInternalIP(net.ParseIP(ip)) {
			return ip
		}
	}

	if ip = strings.TrimSpace(req.Header.Get("X-Real-Ip")); ip != "" && !IsInternalIP(net.ParseIP(ip)) {
		return ip
	}

	if ip, _, _ = net.SplitHostPort(req.RemoteAddr); !IsInternalIP(net.ParseIP(ip)) {
		return ip
	}

	return ip
}

// GetMacAddresses get mac address.
func GetMacAddresses() []string {
	var macAddresses []string

	nets, err := net.Interfaces()
	if err != nil {
		return macAddresses
	}

	for _, v := range nets {
		macAddr := v.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddresses = append(macAddresses, macAddr)
	}

	return macAddresses
}

// GetIps return all ipv4 of system.
// Play: https://go.dev/play/p/NUFfcEmukx1
func GetIps() []string {
	var ips []string

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, addr := range addresses {
		ipNet, isValid := addr.(*net.IPNet)
		if isValid && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
