package vincode

import (
	"golib/tools/vin-code/mfrs/general"
	"golib/tools/vin-code/misc"
)

// note this file not for GM(通用汽车)
type GeneralVINCode string

func (str GeneralVINCode) ParseWMI() (WMIData, error) {
	wmi := WMIData{}
	wmi.Continent = misc.GetVINContinent(string(str))
	wmi.Country = misc.GetVINCountry(string(str))
	wmi.Manufacturer = misc.GetVINManuf(string(str))

	return wmi, nil
}

func (str GeneralVINCode) ParseVDS() (VDSData, error) {
	vds := VDSData{}
	return vds, nil
}

func (str GeneralVINCode) ParseVIS() (VISData, error) {
	re := general.GetVISRune(string(str))
	vis := VISData{}
	vis.SequenceNO = re.SequenceNO
	vis.ModelYear = misc.GetModelYearStr(re.YearRune)
	if vis.ModelYear == "0" {
		return vis, VINCodeParseYearError
	}

	return vis, nil
}
