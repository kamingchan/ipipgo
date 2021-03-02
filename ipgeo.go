package ipipgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"github.com/tidwall/gjson"
)

var (
	ErrInvalidIP = errors.New("invalid IP address")
	ErrNetwork   = errors.New("network error")
)

type IPGeo struct {
	IP *net.IP

	Country     string
	CountryCode string
	Region      string
	City        string

	ISP string
	ASN int

	Lat float64
	Lon float64
}

func (geo *IPGeo) String() string {
	var ls []string
	if geo.Country != "" {
		ls = append(ls, geo.Country)
	}
	if geo.Region != "" {
		ls = append(ls, geo.Region)
	}
	if geo.City != "" {
		ls = append(ls, geo.City)
	}
	if geo.ISP != "" {
		ls = append(ls, geo.ISP)
	}
	return strings.Join(ls, "，")
}

func GetHostIP() (net.IP, error) {
	resp, err := httpGet("https://api.ip.sb/ip")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
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
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, ErrNetwork
	}
	geo := new(IPGeo)
	// parse asn
	if asn := gjson.GetBytes(b, "asn"); asn.Exists() {
		geo.ASN = int(asn.Int())
	}
	// parse lon, lat
	if latitude := gjson.GetBytes(b, "latitude"); latitude.Exists() {
		geo.Lat = latitude.Float()
	}
	if longitude := gjson.GetBytes(b, "longitude"); longitude.Exists() {
		geo.Lon = longitude.Float()
	}
	// geo
	if country := gjson.GetBytes(b, "country"); country.Exists() {
		geo.Country = country.String()
	}
	if countryCode := gjson.GetBytes(b, "country_code"); countryCode.Exists() {
		geo.CountryCode = countryCode.String()
	}
	if region := gjson.GetBytes(b, "region"); region.Exists() {
		geo.Region = region.String()
	}
	if city := gjson.GetBytes(b, "city"); city.Exists() {
		geo.City = city.String()
	}
	// isp
	if isp := gjson.GetBytes(b, "isp"); isp.Exists() {
		geo.ISP = isp.String()
	}
	return geo, nil
}
