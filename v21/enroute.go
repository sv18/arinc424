package arinc

// The Enroute Airways file will contain the sequential listing of officially published airways and other established ATS Routes by geographical areas.
// Also contains published airways specific to helicopter operations.

// Enroute Primary Record structure
// refer: section 4.1.6.1 Airport Primary Records
type Enroute struct {
	RecordType              string `fixed:"1,1"`   //1 (1)
	CustomerAreaCode        string `fixed:"2,4"`   //2 thru 4 (3)
	SectionCode             string `fixed:"5,5"`   //5 (1)
	SubsectionCode          string `fixed:"6,6"`   //6 (1)
	RouteIdentifier         string `fixed:"14,18"` // 14 thru 18  (5)
	SequenceNumber          string `fixed:"26,29"` // 26 thru 29 (4)
	FixIdentifier           string `fixed:"30,34"` // 30 thru 34 (5)
	ICAOCode                string `fixed:"35,36"` // 35 thru 36 (2)
	SectionCode2            string `fixed:"37,37"` // 37 (1)
	SubSectionCode2         string `fixed:"38,38"` // 38 (1)
	ContinuationRecordNo    string `fixed:"39,39"` // 39 (1)
	WaypointDescriptionCode string `fixed:"40,43"` // 40 thru 43 (4)
	//TODO: add the below fields
	// Boundary Code // 44 (1)
	// Route Type // 45 (1)
	// Level // 46 (1)
	// Direction Restriction // 47 (1)
	// Cruise Table Indicator // 48 thru 49 (2)
	// EU Indicator // 50 (1)
	// Recommended NAVAID // 51 thru 54 (4)
	// ICAO Code // 55 thru 56 (2)
	// RNP // 57 thru 59 (3)
	// Blank (Spacing) // 60 thru 62 (3)
	// Theta // 63 thru 66 (4)
	// Rho // 67 thru 70 (4)
	// Outbound Magnetic Course // 71 thru 74 (4)
	// Route Distance From // 75 thru 78 (4)
	// Inbound Magnetic Course // 79 thru 82 (4)
	// Blank (Spacing) // 83 (1)
	// Minimum Altitude // 84 thru 88 (5)
	// Minimum Altitude // 89 thru 93 (5)
	// Maximum Altitude // 94 thru 98 (5)
	// Fix Radius Transition Indicator // 99 thru 101 (3)
	// Vertical Scale Factor // 102 thru 104 (3)
	// RVSM Minimum Level // 105 thru 107 (3)
	// VSF RVSM Maximum Level // 108 thru 110 (3)
	// Reserved // 111 thru 114 (4)
	// Blank (Spacing) // 115 thru 123 (9)
	// File Record No // 124 thru 128 (5)
	// Cycle Date // 129 thru 132 (4)
}

func LoadEnroute(datFilePath string) ([]Enroute, error) {
	var temp []Enroute
	err := load(datFilePath, &temp)
	if err != nil {
		return nil, err
	}

	var a []Enroute
	for _, r := range temp {
		if r.SectionCode == sectionCodeEnroute { //TODO: add the correct logic to filter the airports
			a = append(a, r)
		}
	}

	return a, nil
}
