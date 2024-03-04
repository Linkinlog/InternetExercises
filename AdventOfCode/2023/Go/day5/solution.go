package main

func part1(input []string) any {
	a := parse(input)
	score := 0

	for _, seed := range a.seeds {
		tScore := a.locationForSeed(seed)
		if score == 0 {
			score = tScore
			continue
		}

		if tScore < score {
			score = tScore
		}
	}

	return score
}

func part2(input []string) any {
	a := parse(input)
	score := 0

	for i := 0; i < len(a.seeds); i += 2 {
		for j := a.seeds[i]; j < a.seeds[i]+a.seeds[i+1]; j++ {
			tScore := a.locationForSeed(j)
			if score == 0 {
				score = tScore
				continue
			}

			if tScore < score {
				score = tScore
			}
		}
	}

	return score
}

const (
	seedSoil   string = "seed-to-soil"
	soilFert   string = "soil-to-fertilizer"
	fertWater  string = "fertilizer-to-water"
	waterLight string = "water-to-light"
	lightTemp  string = "light-to-temperature"
	tempHum    string = "temperature-to-humidity"
	humLoc     string = "humidity-to-location"
)

type almanac struct {
	seeds            []int
	seedToSoilMaps   []conversionMap
	soilToFertMaps   []conversionMap
	fertToWaterMaps  []conversionMap
	waterToLightMaps []conversionMap
	lightToTempMaps  []conversionMap
	tempToHumMaps    []conversionMap
	HumToLocMaps     []conversionMap
}

type conversionMap struct {
	destination int
	source      int
	width       int
}

func (a *almanac) convert(source int, cMap *[]conversionMap) int {
	for _, sMap := range *cMap {
		if source >= sMap.source && source < sMap.source+sMap.width {
			return source + (sMap.destination - sMap.source)
		}
	}
	return source
}

func (a *almanac) locationForSeed(seed int) int {
	soil := a.convert(seed, &a.seedToSoilMaps)
	fert := a.convert(soil, &a.soilToFertMaps)
	water := a.convert(fert, &a.fertToWaterMaps)
	light := a.convert(water, &a.waterToLightMaps)
	temp := a.convert(light, &a.lightToTempMaps)
	hum := a.convert(temp, &a.tempToHumMaps)
	return a.convert(hum, &a.HumToLocMaps)
}
