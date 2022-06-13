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
		host := os.Args[1]
		ips, err := net.LookupIP(host)
		must(err)
		ip = ips[0].String()
		fmt.Printf("NS: %s -> %s\n", host, ip)
	}
	ip = strings.TrimSpace(ip)
	if net.ParseIP(ip) == nil {
		must(fmt.Errorf("%s is not a valid IP address", ip))
	}
	geo, err := ipipgo.GetGeo(ip)
	must(err)
	fmt.Printf("IP: %s\n", ip)
	fmt.Printf("GEO: %s\n", geo)
	fmt.Printf("LOC: %.4f / %.4f -> https://www.openstreetmap.org/#map=12/%f/%f\n", geo.Latitude, geo.Longitude, geo.Latitude, geo.Longitude)
	fmt.Printf("ASN: AS%d -> https://ipinfo.io/AS%d\n", geo.Asn, geo.Asn)
}
