package arinc

// This file contains runway information.

// Runway Primary Record structure
// refer: section 4.1.10.1 Runway Primary Records
type Runway struct {
	RecordType                 string `fixed:"1,1"`     //1(1)
	CustomerAreaCode           string `fixed:"2,4"`     //2 thru 4 (3)
	SectionCode                string `fixed:"5,5"`     //5 (1)
	AirportICAOIdentifier      string `fixed:"7,10"`    //7 thru 10(4)
	ICAOCode                   string `fixed:"11,12"`   //11 thru 12(2)
	SubsectionCode             string `fixed:"13,13"`   //13(1)
	RunwayIdentifier           string `fixed:"14,18"`   //14 thru 18(5)
	ContinuationRecordNo       string `fixed:"22,22"`   //22(1)
	RunwayLength               string `fixed:"23,27"`   //23 thru 27(5)
	RunwayMagneticBearing      string `fixed:"28,31"`   //28 thru 31(4)
	RunwayLatitude             string `fixed:"33,41"`   //33 thru 41(9)
	RunwayLongitude            string `fixed:"42,51"`   //42 thru 51(10)
	RunwayGradient             string `fixed:"52,56"`   //52 thru 56(5)
	EllipsoidHeight            string `fixed:"61,66"`   //61 thru 66(6)
	LandingThresholdElevation  string `fixed:"67,71"`   //67 thru 71(5)
	DisplacedThresholdDistance string `fixed:"72,75"`   //72 thru 75(4)
	RunwayWidth                string `fixed:"78,80"`   //78 thru 80(3)
	TCHValueIndicator          string `fixed:"81,81"`   //81(1)
	Stopway                    string `fixed:"87,90"`   //87 thru 90(4)
	ThresholdCrossingHeight    string `fixed:"96,98"`   //96 thru 98(3)
	RunwayDescription          string `fixed:"102,123"` //102 thru 123(22)
	FileRecordNo               string `fixed:"124,128"` //124 thru 128(5)
	CycleDate                  string `fixed:"129,132"` //129 thru 132(4)
}

func LoadRunways(datFilePath string) ([]Runway, error) {
	var temp []Runway
	err := load(datFilePath, &temp)
	if err != nil {
		return nil, err
	}

	var a []Runway
	for _, r := range temp {
		if r.SectionCode == sectionCodeAirport && r.SubsectionCode == subSectionCodeRunways { 
			a = append(a, r)
		}
	}

	return a, nil
}
