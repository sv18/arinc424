package arinc

// This file contains Waypoint information.

// Waypoint Primary Records structure
// Waypoint Record (EA) or (PC)
// refer: section 4.1.4.1 Waypoint Primary Records
/*
The Enroute Waypoint file (EA) contains all enroute on-airway and off-airway waypoints within a desired geographical area.
The Airport Terminal Waypoint file (PC) contains all terminal waypoints and VFR waypoints within the geographical area of each airport. Airport Terminal Waypoints utilized by two or more airports will be stored in the Enroute Waypoint Subsection (EA) to eliminate duplication.
Terminal Waypoints used jointly by an airport and a heliport are also stored in the Enroute Waypoint file.
The Enroute Waypoint File will contain waypoints established for Helicopter Airways.
For Heliport Terminal Waypoints (HC), see Section 4.2.2.
*/
type Waypoint struct {
	RecordType               string `fixed:"1,1"`     //(1)1
	CustomerAreaCode         string `fixed:"2,4"`     //(3)2 thru 4
	SectionCode              string `fixed:"5,5"`     //(1)5
	SubsectionCode           string `fixed:"6,6"`     //(1)6
	RegionCode               string `fixed:"7,10"`    //(4)7 thru 10
	ICAOCode                 string `fixed:"11,12"`   //(2)11 thru 12
	Subsection               string `fixed:"13,13"`   //(1)13
	WaypointIdentifier       string `fixed:"14,18"`   //(5)14 thru 18
	ICAOCode2                string `fixed:"20,21"`   //(2)20 thru 21
	ContinuationRecordNo     string `fixed:"22,22"`   //(1)22
	WaypointType             string `fixed:"27,29"`   //(3)27 thru 29
	WaypointUsage            string `fixed:"31,31"`   //(1)31
	WaypointLatitude         string `fixed:"33,41"`   //(9)33 thru 41
	WaypointLongitude        string `fixed:"42,51"`   //(10)42 thru 51
	DynamicMagneticVariation string `fixed:"75,79"`   //(5)75 thru 79
	DatumCode                string `fixed:"85,87"`   //(3)85 thru 87
	NameFormatIndicator      string `fixed:"96,98"`   //(3)96 thru 98
	WaypointNameDescription  string `fixed:"99,123"`  //(25)99 thru 123
	FileRecordNo             string `fixed:"124,128"` //(5)124 thru 128
	CycleDate                string `fixed:"129,132"` //(4)129 thru 132
}

func LoadWaypoints(datFilePath string) ([]Waypoint, error) {
	var temp []Waypoint
	err := load(datFilePath, &temp)
	if err != nil {
		return nil, err
	}

	var a []Waypoint
	for _, r := range temp {
		if r.SectionCode == sectionCodeEnroute && r.SubsectionCode == subSectionCodeWaypoints ||
			r.SectionCode == sectionCodeAirport && r.SubsectionCode == subSectionCodeTerminalWaypoints {
			a = append(a, r)
		}
	}

	return a, nil
}
