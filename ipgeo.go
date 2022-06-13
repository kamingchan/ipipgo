package ipipgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	emojiflag "github.com/jayco/go-emoji-flag"
)

var (
	ErrInvalidIP = errors.New("invalid IP address")
)

type IPGeo struct {
	Organization    string  `json:"organization"`
	Longitude       float64 `json:"longitude"`
	City            string  `json:"city"`
	Timezone        string  `json:"timezone"`
	Isp             string  `json:"isp"`
	Offset          int     `json:"offset"`
	Region          string  `json:"region"`
	Asn             int     `json:"asn"`
	AsnOrganization string  `json:"asn_organization"`
	Country         string  `json:"country"`
	Ip              string  `json:"ip"`
	Latitude        float64 `json:"latitude"`
	ContinentCode   string  `json:"continent_code"`
	CountryCode     string  `json:"country_code"`
	RegionCode      string  `json:"region_code"`
}

func (geo *IPGeo) String() string {
	var ls []string
	if geo.Country != "" {
		ls = append(ls, geo.Country+" "+emojiflag.GetFlag(geo.CountryCode))
	}
	if geo.Region != "" {
		ls = append(ls, geo.Region)
	}
	if geo.City != "" {
		ls = append(ls, geo.City)
	}
	if geo.AsnOrganization != "" {
		ls = append(ls, geo.AsnOrganization)
	}
	return strings.Join(ls, " / ")
}

func GetHostIP() (net.IP, error) {
	resp, err := httpGet("https://api.ip.sb/ip")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ipStr := strings.TrimSpace(string(body))
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, ErrInvalidIP
	}
	return ip, nil
}

func GetGeo(ipStr string) (*IPGeo, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, ErrInvalidIP
	}
	url := fmt.Sprintf("https://api.ip.sb/geoip/%s", ipStr)
	res, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	geo := new(IPGeo)
	err = json.NewDecoder(res.Body).Decode(geo)
	if err != nil {
		return nil, err
	}
	return geo, nil
}
