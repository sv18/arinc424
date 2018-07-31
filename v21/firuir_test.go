package arinc

import (
	"fmt"
	"testing"
)

func TestFirUir(t *testing.T) {

	a, err := LoadFirUir("arincdata-full.dat")
	if err != nil {
		t.Error(err)
		return
	}

	dup := map[string]int{}
	for _, r := range a {
		//fmt.Printf("%#v\n", r)
		dup[r.FirUirName]++
	}
	fmt.Println("FirUirName::::::::::::::::", len(a))
	for k, v := range dup {
		if v > 1 {
			fmt.Println(k, v)
		}
	}

}
