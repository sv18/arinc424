package arinc

// Airport Primary Record structure
// refer: section 4.1.7.1 Airport Primary Records
type Airport struct {
	RecordType                     string `fixed:"1,1"`
	CustomerAreaCode               string `fixed:"2,4"`
	SectionCode                    string `fixed:"5,5"`
	AirportICAOIdentifier          string `fixed:"7,10"`  //7 thru 10 (4)
	ICAOCode                       string `fixed:"11,12"` //11 thru 12 (2)
	SubsectionCode                 string `fixed:"13,13"`
	AtaIataDesignator              string `fixed:"14,16"`   //14 thru 16 (3)
	ContinuationRecordNo           string `fixed:"22,22"`   // 22 (1)
	SpeedLimitAltitude             string `fixed:"23,27"`   // 23 thru 27 (5)
	LongestRunway                  string `fixed:"28,30"`   //28 thru 30 (3)
	IFRCapability                  string `fixed:"31,31"`   // 31 (1)
	LongestRunwaySurfaceCode       string `fixed:"32,32"`   //32 (1)
	AirportReferencePointLatitude  string `fixed:"33,41"`   //33 thru 41(9)
	AirportReferencePointLongitude string `fixed:"42,51"`   // 42 thru 51(10)
	MagneticVariation              string `fixed:"52,56"`   //52 thru 56 (5)
	AirportElevation               string `fixed:"57,61"`   // 57 thru 61(5)
	SpeedLimit                     string `fixed:"62,64"`   // 62 thru 64(3)
	RecommendedNavaid              string `fixed:"65,68"`   // 65 thru 68(4)
	ICAOCode2                      string `fixed:"69,70"`   // 69 thru 70 (2)
	TransitionsAltitude            string `fixed:"71,75"`   // 71 thru 75 (5)
	TransitionLevel                string `fixed:"76,80"`   // 76 thru 80 (5)
	PublicMilitaryIndicator        string `fixed:"81,81"`   // 81  (1)
	TimeZone                       string `fixed:"82,84"`   // 82 thru 84 (3)
	DaylightIndicator              string `fixed:"85,85"`   // 85 (1)
	MagneticTrueIndicator          string `fixed:"86,86"`   // 86 (1)
	DatumCode                      string `fixed:"87,89"`   // 87 thru 89 (3)
	AirportName                    string `fixed:"94,123"`  // 94 thru 123 (30)
	FileRecordNumber               string `fixed:"124,128"` // 124 thru 128 (5)
	CycleDate                      string `fixed:"129,132"` // 129 thru 132 (4)
}

func LoadAirport(datFilePath string) ([]Airport, error) {

	var temp []Airport
	err := load(datFilePath, &temp)
	if err != nil {
		return nil, err
	}

	var a []Airport
	for _, r := range temp {
		if r.SectionCode == sectionCodeAirport { //TODO: add the correct logic to filter the airports. Look like we nned to filter subcode = A filter
			a = append(a, r)
		}
	}

	return a, nil
}
