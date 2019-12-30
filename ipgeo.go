package ipipgo

import (
	"encoding/json"
	"errors"
	"net"
)

var (
	ErrInvalidIP = errors.New("invalid IP address")
	ErrDecode    = errors.New("json decode failed")
)

const (
	ipipResponseLen = 5
)

type IPGeo struct {
	IP *net.IP

	Country  string
	Province string
	City     string

	ISP string
}

func GetGeo(ip string) (*IPGeo, error) {
	_ip := net.ParseIP(ip)
	if _ip == nil {
		return nil, ErrInvalidIP
	}
	res, err := httpGet("http://freeapi.ipip.net/" + ip)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var resp []string
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, ErrDecode
	}
	if len(resp) != ipipResponseLen {
		return nil, ErrDecode
	}
	return &IPGeo{
		IP:       &_ip,
		Country:  resp[0],
		Province: resp[1],
		City:     resp[2],
		ISP:      resp[4],
	}, nil
}
