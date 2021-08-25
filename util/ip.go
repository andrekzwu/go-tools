package util

import "net"

/**
fget local ip
*/
func GetLocalIP() string {
    addrs, _ := net.InterfaceAddrs()
    for _, address := range addrs {
        // check the ip is loop address?
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }

        }
    }
    return "127.0.0.1"
}
