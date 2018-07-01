package arinc

import (
	"io/ioutil"
	"os"

	fixedwidth "github.com/ianlopshire/go-fixedwidth"
)

// Airport Primary Record structure
// refer: section 4.1.7.1 Airport Primary Records
type Airport struct {
	RecordType                     string `fixed:"1,1"`
	CustomerAreaCode               string `fixed:"2,4"`
	SectionCode                    string `fixed:"5,5"`
	AirportIcaoIdentifier          string //7 thru 10 (4)
	ICAOCode                       string //11 thru 12 (2)
	SubSecCode                     string `fixed:"13,13"`
	AtaIataDesignator              string //14 thru 16 (3)
	ContinuationRecordNumber       string // 22 (1)
	SpeedLimitAltitude             string // 23 thru 27 (5)
	LongestRunway                  string //28 thru 30 (3)
	IFRCapability                  string // 31 (1)
	LongestRunwaySurfaceCode       string //32 (1)
	AirportReferencePointLatitude  string //33 thru 41(9)
	AirportReferencePointLongitude string // 42 thru 51(10)
	MagneticVariation              string //52 thru 56 (5)
	AirportElevation               string // 57 thru 61(5)
	SpeedLimit                     string // 62 thru 64(3)
	RecommendedNavaid              string // 65 thru 68(4)
	ICAOCode2                      string // 69 thru 70 (2)
	TransitionsAltitude            string // 71 thru 75 (5)
	TransitionLevel                string // 76 thru 80 (5)
	PublicMilitaryIndicator        string // 81  (1)
	TimeZone                       string // 82 thru 84 (3)
	DaylightIndicator              string // 85 (1)
	MagneticTrueIndicator          string // 86 (1)
	DatumCode                      string // 87 thru 89 (3)
	AirportName                    string // 94 thru 123 (30)
	FileRecordNumber               string // 124 thru 128 (5)
	CycleDate                      string // 129 thru 132 (4)
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
