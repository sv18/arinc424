// this package implements as per the ARINC 424 Specifications version 21

package arinc

import (
	"io/ioutil"
	"os"

	fixedwidth "github.com/ianlopshire/go-fixedwidth"
)

const ( // record type codes: ref specs section 5.2
	recordTypeStandard = "S"
	recordTypeTailored = "T"
)

// section & subsection codes: ref specs table 5-1
const (
	sectionCodeMORA        = "A"
	subSectionCodeGridMORA = "S"

	sectionCodeNavaid             = "D"
	subSectionCodeVHFNavaid       = ""
	subSectionCodeNDBNavaid       = "B"
	subSectionCodeTACANDuplicates = "T"

	sectionCodeEnroute                 = "E"
	subSectionCodeWaypoints            = "A"
	subSectionCodeAirwayMarkers        = "M"
	subSectionCodeHoldingPatterns      = "P"
	subSectionCodeAirwaysAndRoutes     = "R"
	subSectionCodeSpecialActivityAreas = "S"
	subSectionCodePreferredRoutes      = "T"
	subSectionCodeAirwayRestrictions   = "U"
	subSectionCodeCommunications       = "V"

	sectionCodeHeliport              = "H"
	subSectionCodeReferencePoints    = "A"
	subSectionCodeTerminalWaypoints  = "C"
	subSectionCodeSIDs               = "D"
	subSectionCodeSTARs              = "E"
	subSectionCodeApproachProcedures = "F"
	subSectionCodeHelipads           = "H"
	subSectionCodeTAA                = "K"
	subSectionCodeMSA                = "S"
	subSectionCodeSBASPathPoint      = "P"
	//subSectionCodeCommunications = "V" // same as under other section

	sectionCodeAirport = "P"
	//subSectionCodeReferencePoints = "A" // same as under other section
	subSectionCodeGates = "B"
	//subSectionCodeTerminalWaypoints  = "C" // same as under other section
	//subSectionCodeSIDs = "D" // same as under other section
	//subSectionCodeSTARs = "E" // same as under other section
	//subSectionCodeApproachProcedures = "F" // same as under other section
	subSectionCodeRunways = "G"
	//subSectionCodeHelipads = "H" // same as under other section
	subSectionCodeLocalizerGlideslope = "I"
	//subSectionCodeTAA = "K" // same as under other section
	subSectionCodeMLS             = "L"
	subSectionCodeLocalizerMarker = "M"
	subSectionCodeTerminalNDB     = "N"
	//subSectionCodeSBASPathPoint = "P" // same as under other section
	subSectionCodeGBASPathPoint     = "Q"
	subSectionCodeFltPlanningArrDep = "R"
	//subSectionCodeMSA = "S" // same as under other section
	subSectionCodeGLSStation = "T"
	//subSectionCodeCommunications = "V" // same as under other section

	sectionCodeCompanyRoutes                = "R"
	subSectionCodeCompanyRoutes             = "" // (Master Airline File)
	subSectionCodeAlternateRecords          = "A"
	subSectionCodeHelicopterOperationRoutes = "H" // (Master Helicopter File)

	sectionCodeTables                   = "T"
	subSectionCodeCruisingTables        = "C"
	subSectionCodeGeographicalReference = "G"
	subSectionCodeRNAVNameTable         = "N"
	subSectionCodeCommunicationType     = "V"

	sectionCodeAirspace               = "U"
	subSectionCodeControlledAirspace  = "C"
	subSectionCodeFIR_UIR             = "F"
	subSectionCodeRestrictiveAirspace = "R"
)

func load(datFilePath string, out interface{}) error {

	f, err := os.Open(datFilePath) //TODO: handle error
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(f) //TODO: handle error
	if err != nil {
		return err
	}

	err = fixedwidth.Unmarshal(b, out)
	if err != nil {
		return err
	}

	return nil
}
