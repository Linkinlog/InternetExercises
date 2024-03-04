package main

import (
	"strconv"
	"strings"
)

func parse(s []string) almanac {
	seeds := parseSeederoonies(s[0])

	stringToConversion := map[string][]conversionMap{
		seedSoil:   {},
		soilFert:   {},
		fertWater:  {},
		waterLight: {},
		lightTemp:  {},
		tempHum:    {},
		humLoc:     {},
	}

	for i := 1; i < len(s); i++ {
		if strings.TrimSpace(s[i]) == "" {
			continue
		}

		for v := range stringToConversion {
			if i < len(s) && strings.Contains(s[i], v) {
				i += 1 // go to actual map
				for i < len(s) && strings.TrimSpace(s[i]) != "" {
					if _, ok := stringToConversion[v]; !ok {
						panic("woah, that was weird.")
					}
					stringToConversion[v] = append(stringToConversion[v], parseConversionMap(s[i]))
					i += 1 // next map
				}
			}
		}
	}

	return almanac{
		seeds:            seeds,
		seedToSoilMaps:   stringToConversion[seedSoil],
		soilToFertMaps:   stringToConversion[soilFert],
		fertToWaterMaps:  stringToConversion[fertWater],
		waterToLightMaps: stringToConversion[waterLight],
		lightToTempMaps:  stringToConversion[lightTemp],
		tempToHumMaps:    stringToConversion[tempHum],
		HumToLocMaps:     stringToConversion[humLoc],
	}
}

func parseSeederoonies(s string) []int {
	halves := strings.Split(s, ": ")
	inties := []int{}
	for _, hopefullyANumber := range strings.Split(halves[1], " ") {
		number, _ := strconv.Atoi(hopefullyANumber)
		inties = append(inties, number)
	}

	return inties
}

func parseConversionMap(s string) conversionMap {
	splitsVille := strings.Split(s, " ")
	destination, _ := strconv.Atoi(splitsVille[0])
	source, _ := strconv.Atoi(splitsVille[1])
	width, _ := strconv.Atoi(splitsVille[2])

	return conversionMap{
		destination: destination,
		source:      source,
		width:       width,
	}
}
