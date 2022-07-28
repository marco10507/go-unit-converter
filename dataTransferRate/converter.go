// Use this package to convert data transfer rates
package dataTransferRate

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	Kbps = "Kbps"
	Mbps = "Mbps"
	Gbps = "Gbps"
	Tbps = "Tbps"
)

const (
	kb = 1000 // 1KB is 1000 bytes
	mb = kb * 1000
	gb = mb * 1000
	tb = gb * 1000
)

var units = createUnitRegistry()

type unitRegistry struct {
	Kbps  string
	Mbps  string
	Gbps  string
	Tbps  string
	units []string
}

func (u *unitRegistry) List() []string {
	return u.units
}

func (u *unitRegistry) Parse(unit string) (string, error) {
	for _, v := range u.units {
		if strings.ToLower(v) == strings.ToLower(unit) {
			return v, nil
		}
	}

	return "", fmt.Errorf("The unit %s must be one of the following: %s", unit, u.units)
}

func createUnitRegistry() unitRegistry {
	registry := unitRegistry{
		Kbps:  Kbps,
		Mbps:  Mbps,
		Gbps:  Gbps,
		Tbps:  Tbps,
		units: []string{Kbps, Mbps, Gbps, Tbps},
	}

	return registry
}

type dataTransferRate struct {
	value float64
	unit  string // Kbps, Mbps, etc
}

func removeAllWhiteSpaces(input string) string {
	pattern := `\s+`
	r, err := regexp.Compile(pattern)

	if err != nil {
		return ""
	}

	output := r.ReplaceAllString(input, "")

	return output
}

func validateTransferRateParts(matches []string) error {
	if len(matches) != 2 {
		return errors.New("Transfer rate must " +
			"be compose of two parts value and units. For example: 100 Mbps")
	}

	value := matches[0]
	unit := matches[1]

	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return fmt.Errorf("The value %s must be a number such as: 10, 10.5, etc", value)
	}

	_, err := units.Parse(unit)

	if err != nil {
		return err
	}

	return nil
}

func parseToDataTransferRate(rawData string) (dataTransferRate, error) {
	pattern := `(?m)([\d\.]+)|(\D+)`
	r, err := regexp.Compile(pattern)

	if err != nil {
		return dataTransferRate{}, err
	}

	rawData = removeAllWhiteSpaces(rawData)
	matches := r.FindAllString(rawData, -1)

	err = validateTransferRateParts(matches)

	if err != nil {
		return dataTransferRate{}, err
	}

	floatValue, _ := strconv.ParseFloat(matches[0], 64)

	output := dataTransferRate{
		value: floatValue,
		unit:  matches[1],
	}

	return output, nil
}

func convertToBytes(fromTransferRate dataTransferRate) (float64, error) {
	var bytes float64

	unit, err := units.Parse(fromTransferRate.unit)
	if err != nil {
		return 0, err
	}

	value := fromTransferRate.value

	switch unit {
	case Kbps:
		bytes = value * kb
	case Mbps:
		bytes = value * mb
	case Gbps:
		bytes = value * gb
	case Tbps:
		bytes = value * tb
	default:
		return 0, fmt.Errorf("Unit %s not found", unit)
	}

	return bytes, nil
}

type converter struct {
	transferRate string
}

func (c *converter) Convert(dataTransferRate string) *converter {
	c.transferRate = dataTransferRate

	return c
}

func (c *converter) To(toUnit string) (float64, error) {
	transferRate, err := parseToDataTransferRate(c.transferRate)
	if err != nil {
		return 0, err
	}

	unit, err := units.Parse(toUnit)
	if err != nil {
		return 0, err
	}

	bytes, err := convertToBytes(transferRate)
	if err != nil {
		return 0, err
	}

	var convertedVale float64

	switch unit {
	case Kbps:
		convertedVale = bytes / kb
	case Mbps:
		convertedVale = bytes / mb
	case Gbps:
		convertedVale = bytes / gb
	case Tbps:
		convertedVale = bytes / tb
	default:
		return 0, fmt.Errorf("Unit %s not found", unit)
	}

	return convertedVale, nil
}

// NewConverter gets a data transfer rate converter
func NewConverter() *converter {
	return &converter{
		transferRate: "",
	}
}
