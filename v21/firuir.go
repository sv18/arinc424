package arinc

// This file contains FIR/UIR information.

// FirUir is the FIR/UIR Record structure
// refer: section 4.1.17.1 FIR/UIR Primary Records
type FirUir struct {
	RecordType             string `fixed:"1,1"`     //(1)	1
	CustomerAreaCode       string `fixed:"2,4"`     //(3)	2 thru 4
	SectionCode            string `fixed:"5,5"`     //(1)	5
	SubsectionCode         string `fixed:"6,6"`     //(1)	6
	FirUirIdentifier       string `fixed:"7,10"`    //(4)	7 thru 10
	FirUirAddress          string `fixed:"11,14"`   //(4)	11 thru 14
	FirUirIndicator        string `fixed:"15,15"`   //(1)	15
	SequenceNumber         string `fixed:"16,19"`   //(4)	16 thru 19
	ContinuationRecordNo   string `fixed:"20,20"`   //(1)	20
	AdjacentFirIdentifier  string `fixed:"21,24"`   //(4)	21 thru 24
	AdjacentUirIdentifier  string `fixed:"25,28"`   //(4)	25 thru 28
	ReportingUnitsSpeed    string `fixed:"29,29"`   //(1)	29
	ReportingUnitsAltitude string `fixed:"30,30"`   //(1)	30
	EntryReport            string `fixed:"31,31"`   //(1)	31
	BoundaryVia            string `fixed:"33,34"`   //(2)	33 thru 34
	FirUirLatitude         string `fixed:"35,43"`   //(9)	35 thru 43
	FirUirLongitude        string `fixed:"44,53"`   //(10)	44 thru 53
	ArcOriginLatitude      string `fixed:"54,62"`   //(9)	54 thru 62
	ArcOriginLongitude     string `fixed:"63,72"`   //(10)	63 thru 72
	ArcDistance            string `fixed:"73,76"`   //(4)	73 thru 76
	ArcBearing             string `fixed:"77,80"`   //(4)	77 thru 80
	FirUpperLimit          string `fixed:"81,85"`   //(5)	81 thru 85
	UirLowerLimit          string `fixed:"86,90"`   //(5)	86 thru 90
	UirUpperLimit          string `fixed:"91,95"`   //(5)	91 thru 95
	CruiseTableInd         string `fixed:"96,97"`   //(2)	96 thru 97
	FirUirName             string `fixed:"99,123"`  //(25)	99 thru 123
	FileRecordNo           string `fixed:"124,128"` //(5)	124 thru 128
	CycleDate              string `fixed:"129,132"` //(4)	129 thru 132
}

func LoadFirUir(datFilePath string) ([]FirUir, error) {
	var temp []FirUir
	err := load(datFilePath, &temp)
	if err != nil {
		return nil, err
	}

	var a []FirUir
	for _, r := range temp {
		if r.SectionCode == sectionCodeAirspace &&
			r.SubsectionCode == subSectionCodeFIR_UIR {
			a = append(a, r)
		}
	}

	return a, nil
}
