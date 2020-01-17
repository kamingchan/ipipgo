package ipipgo

import (
	"fmt"
	"testing"
)

func TestGetGeo(t *testing.T) {
	geo, err := GetGeo("218.107.55.254")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", geo)
}
