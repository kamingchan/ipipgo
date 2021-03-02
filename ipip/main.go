package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/kamingchan/ipipgo/v2"
)

func must(err error) {
	if err != nil {
		fmt.Printf("encounter error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	var ip string
	if len(os.Args) == 1 {
		_ip, err := ipipgo.GetHostIP()
		must(err)
		ip = _ip.String()
	} else {
		ip = os.Args[1]
	}
	ip = strings.TrimSpace(ip)
	if net.ParseIP(ip) == nil {
		must(fmt.Errorf("%s is not a valid IP address", ip))
	}
	geo, err := ipipgo.GetGeo(ip)
	must(err)
	fmt.Printf("IP: %s\n%v\nAS%d\n", ip, strings.ToUpper(geo.String()), geo.ASN)
}
