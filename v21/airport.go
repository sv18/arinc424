package arinc

import (
	"io/ioutil"
	"os"

	fixedwidth "github.com/ianlopshire/go-fixedwidth"
)

const ( //TODO: move to a common file
	sectionCodeAirport = "P"
)

// Airport Primary Record structure
// refer: section 4.1.7.1 Airport Primary Records
type Airport struct {
	RecordType               string `fixed:"1,1"`
	CustomerAreaCode         string `fixed:"2,4"`
	SectionCode              string `fixed:"5,5"`
	AirportIcaoIdentifier    string //7 thru 10 (4)
	ICAOCode                 string //11 thru 12 (2)
	SubSecCode               string `fixed:"13,13"`
	AtaIataDesignator        string //14 thru 16 (3)
	ContinuationRecordNumber string // 22 (1)
	SpeedLimitAltitude       string // 23 thru 27 (5)
	//TODO more

}

func LoadAirport(datFilePath string) ([]Airport, error) {

	f, err := os.Open(datFilePath) //TODO: handle error
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f) //TODO: handle error
	if err != nil {
		return nil, err
	}

	var temp []Airport
	err = fixedwidth.Unmarshal(b, &temp)
	if err != nil {
		return nil, err
	}

	var a []Airport
	for _, r := range temp {
		if r.SectionCode == sectionCodeAirport { //TODO: add the correct logic to filter the airports
			a = append(a, r)
		}
	}

	return a, nil
}
