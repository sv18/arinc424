package arinc

import (
	"fmt"
	"testing"
)

func TestLoadAirport(t *testing.T) {

	a, err := LoadAirport("arincdata-full.dat")
	if err != nil {
		t.Error(err)
		return
	}

	for _, r := range a {
		fmt.Printf("%#v\n", r)
	}

}
