package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
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
		resp, err := http.Get("https://api.ip.sb/ip")
		must(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		must(err)
		ip = string(body)
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
