package ipipgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var (
	ErrInvalidIP = errors.New("invalid IP address")
	ErrDecode    = errors.New("json decode failed")
)

const (
	responseLen = 7
)

type IPGeo struct {
	IP *net.IP

	Country  string
	Province string
	City     string

	ISP string
	AS  string

	Lat float64
	Lon float64
}

func (geo *IPGeo) String() string {
	var ls []string
	if geo.Country != "" {
		ls = append(ls, geo.Country)
	}
	if geo.Province != "" {
		ls = append(ls, geo.Province)
	}
	if geo.City != "" {
		ls = append(ls, geo.City)
	}
	if geo.ISP != "" {
		ls = append(ls, geo.ISP)
	}
	return strings.Join(ls, "，")
}

func GetGeo(ipStr string) (*IPGeo, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, ErrInvalidIP
	}
	url := fmt.Sprintf("https://btapi.ipip.net/host/info?lang=CN&ip=%s&host", ipStr)
	res, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, ErrDecode
	}
	_as, ok := resp["as"]
	if !ok {
		return nil, ErrDecode
	}
	_area, ok := resp["area"]
	if !ok {
		return nil, ErrDecode
	}
	area := _area.(string)
	ls := strings.Split(area, "\t")
	if len(ls) != responseLen {
		return nil, ErrDecode
	}
	lat, _ := strconv.ParseFloat(ls[5], 64)
	lon, _ := strconv.ParseFloat(ls[6], 64)
	return &IPGeo{
		AS:       _as.(string),
		IP:       &ip,
		Country:  ls[0],
		Province: ls[1],
		City:     ls[2],
		ISP:      ls[4],
		Lat:      lat,
		Lon:      lon,
	}, nil
}
