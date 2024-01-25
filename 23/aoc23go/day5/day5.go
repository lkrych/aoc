package day5

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lkrych/aoc23go/input"
)

func removeWhitespace(s string) string {
	fields := strings.Fields(s) // Fields splits the string s around each instance of one or more consecutive white space characters
	return strings.Join(fields, "")
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

func Part1() {
	// BOILERPLATE for getting file name from stdIn and reading line by line
	filename := flag.String("f", "", "input file")
	// Parse the command-line arguments to read the flag value
	flag.Parse()
	filepath := fmt.Sprintf("../input/%s", *filename)
	scanner, err := input.ReadInputFile(filepath)
	if err != nil {
		panic(err)
	}
	defer scanner.Scan() // Close the file when done reading

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	seeds := []int{}
	seedToSoilMap := map[int]int{}
	soilToFertilizer := map[int]int{}
	fertilizerToWater := map[int]int{}
	waterToLight := map[int]int{}
	lightToTemp := map[int]int{}
	tempToHumidity := map[int]int{}
	humidityToLoc := map[int]int{}

	stringToMap := map[string]map[int]int{
		"seed-to-soil":            seedToSoilMap,
		"soil-to-fertilizer":      soilToFertilizer,
		"fertilizer-to-water":     fertilizerToWater,
		"water-to-light":          waterToLight,
		"light-to-temperature":    lightToTemp,
		"temperature-to-humidity": tempToHumidity,
		"humidity-to-location":    humidityToLoc,
	}

	// strategy, first fill maps, then trace through maps
	var currentMapName string
	hasBeenMapped := map[int]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// load the current working map
		if strings.Contains(line, "seeds") {
			seedSplit := strings.Split(line, " ")
			seedsStr := seedSplit[1:]
			for _, s := range seedsStr {
				seed := convertStringToInt(s)
				seeds = append(seeds, seed)
			}
		} else if strings.Contains(line, "map") {
			splits := strings.Split(line, " ")
			fmt.Println(splits)
			mapName := removeWhitespace(splits[0])
			fmt.Println("Updating map: ", mapName)
			// seed next map with previous map
			for _, s := range seeds {
				prevMap := stringToMap[currentMapName]
				prevVal, ok := prevMap[s]
				if !ok {
					// if a previous value wasn't found default to seed value
					prevVal = s
					fmt.Println("Defaulting to seed value for map ", mapName)
				}
				currentMap := stringToMap[mapName]
				currentMap[s] = prevVal
				stringToMap[mapName] = currentMap
				hasBeenMapped[s] = false
			}
			currentMapName = mapName
		} else if len(removeWhitespace(line)) > 1 {
			if currentMap, ok := stringToMap[currentMapName]; !ok {
				log.Fatalf("Couldn't find %s in map %v", currentMapName, stringToMap)
			} else {
				// we are parsing some ranges in the map
				splitRanges := strings.Split(line, " ")
				destRangeStart := convertStringToInt(splitRanges[0])
				sourceRangeStart := convertStringToInt(splitRanges[1])
				rangeLen := convertStringToInt(splitRanges[2])
				fmt.Println("Parsing ", currentMapName, " dest: ", destRangeStart, ", src: ", sourceRangeStart, ", range: ", rangeLen)
				// iterate through each seed value in currentMap
				for seed, v := range currentMap {
					hasAlreadyBeenMapped := hasBeenMapped[seed]
					if v >= sourceRangeStart && v <= sourceRangeStart+rangeLen && !hasAlreadyBeenMapped {
						// find the difference
						diff := v - sourceRangeStart
						newDest := destRangeStart + diff
						fmt.Println("Assigning ", newDest, " to ", currentMapName, " for seed ", seed)
						currentMap[seed] = newDest
						hasBeenMapped[seed] = true
					}
				}
			}
		}
	}

	lowestLoc := 1000000000000000000
	for _, v := range stringToMap["humidity-to-location"] {
		if v < lowestLoc {
			lowestLoc = v
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Lowest Location", lowestLoc)
}

type SeedRange struct {
	start         int
	len           int
	hasBeenMapped bool
}

func parseLine(line string) (int, int, int) {
	parts := strings.Fields(line)
	return convertStringToInt(parts[0]), convertStringToInt(parts[1]), convertStringToInt(parts[2])
}

func printSeedRangeInfo(sr SeedRange, prevMapName string, sourceRangeStart, rangeLen int) {
	fmt.Printf("Reading %v from %s and comparing to start: %d, len: %d\n", sr, prevMapName, sourceRangeStart, rangeLen)
}

func addRange(ranges []SeedRange, start, len int) []SeedRange {
	return append(ranges, SeedRange{start: start, len: len, hasBeenMapped: false})
}

func handleSeedRange(sr SeedRange, mapStart, mapLen, destRangeStart int) (SeedRange, []SeedRange, []SeedRange) {
	var seedRangesToSave, seedRangesToRead []SeedRange
	srFinish, mapFinish := sr.start+sr.len-1, mapStart+mapLen-1

	switch {
	case sr.start >= mapStart && srFinish <= mapFinish:
		diff := sr.start - mapStart
		seedRangesToSave = addRange(seedRangesToSave, destRangeStart+diff, sr.len)
		sr.hasBeenMapped = true
	case sr.start <= mapStart && srFinish >= mapStart && srFinish <= mapFinish:
		nonOverlapLen := mapStart - sr.start
		overLapLen := srFinish - mapStart + 1
		seedRangesToRead = addRange(seedRangesToRead, sr.start, nonOverlapLen)
		seedRangesToSave = addRange(seedRangesToSave, destRangeStart, overLapLen)
		sr.hasBeenMapped = true
	case sr.start >= mapStart && sr.start <= mapFinish && srFinish >= mapFinish:
		nonOverlapStart := mapFinish + 1
		nonOverlapLen := srFinish - nonOverlapStart + 1
		seedRangesToRead = addRange(seedRangesToRead, nonOverlapStart, nonOverlapLen)
		diff := sr.start - mapStart
		newLen := mapFinish - sr.start + 1
		seedRangesToSave = addRange(seedRangesToSave, destRangeStart+diff, newLen)
		sr.hasBeenMapped = true
	}

	return sr, seedRangesToSave, seedRangesToRead
}

// ProcessLine processes a line of text and updates the maps with seed ranges.
func ProcessLine(line, currentMapName, prevMapName string, stringToMap map[string][]SeedRange) {
	if len(removeWhitespace(line)) <= 1 {
		return
	}

	destRangeStart, sourceRangeStart, rangeLen := parseLine(line)
	processSeedRanges(stringToMap, prevMapName, currentMapName, sourceRangeStart, rangeLen, destRangeStart)
}

// processSeedRanges processes each seed range in the previous map.
func processSeedRanges(stringToMap map[string][]SeedRange, prevMapName, currentMapName string, sourceRangeStart, rangeLen, destRangeStart int) {
	for idx, seedRange := range stringToMap[prevMapName] {
		if seedRange.hasBeenMapped {
			continue
		}

		printSeedRangeInfo(seedRange, prevMapName, sourceRangeStart, rangeLen)
		updatedSeedRange, toSaveRanges, toAppendRanges := handleSeedRange(seedRange, sourceRangeStart, rangeLen, destRangeStart)

		if updatedSeedRange.hasBeenMapped {
			stringToMap[prevMapName][idx] = updatedSeedRange
			stringToMap[currentMapName] = append(stringToMap[currentMapName], toSaveRanges...)
		}

		if len(toAppendRanges) > 0 {
			stringToMap[prevMapName] = append(stringToMap[prevMapName], toAppendRanges...)
		}
	}
}

func Part2() {
	// BOILERPLATE for getting file name from stdIn and reading line by line
	filename := flag.String("f", "", "input file")
	// Parse the command-line arguments to read the flag value
	flag.Parse()
	filepath := fmt.Sprintf("../input/%s", *filename)

	scanner, err := input.ReadInputFile(filepath)
	if err != nil {
		panic(err)
	}
	defer scanner.Scan() // Close the file when done reading

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	initialMap := []SeedRange{}
	seedToSoilMap := []SeedRange{}
	soilToFertilizer := []SeedRange{}
	fertilizerToWater := []SeedRange{}
	waterToLight := []SeedRange{}
	lightToTemp := []SeedRange{}
	tempToHumidity := []SeedRange{}
	humidityToLoc := []SeedRange{}

	stringToMap := map[string][]SeedRange{
		"initial-map":             initialMap,
		"seed-to-soil":            seedToSoilMap,
		"soil-to-fertilizer":      soilToFertilizer,
		"fertilizer-to-water":     fertilizerToWater,
		"water-to-light":          waterToLight,
		"light-to-temperature":    lightToTemp,
		"temperature-to-humidity": tempToHumidity,
		"humidity-to-location":    humidityToLoc,
	}
	// save new ranges to currentMap
	currentMapName := "initial-map"
	// read from old map
	prevMapName := "initial-map"

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// load the current working map
		if strings.Contains(line, "seeds") {
			seedSplit := strings.Split(line, " ")
			seedsStr := seedSplit[1:]
			// 79 14 and  55 13 seeds are now considered a range
			for i := 0; i < len(seedsStr); i += 2 {
				seedSrc := convertStringToInt(seedsStr[i])
				seedRange := convertStringToInt(seedsStr[i+1])
				initialMap = append(initialMap, SeedRange{seedSrc, seedRange, false})
			}
			// fmt.Printf("Initial map: %v\n", initialMap)
			stringToMap["initial-map"] = initialMap
		} else if strings.Contains(line, "map") {
			// if we come across the line map, its time to load the new map
			splits := strings.Split(line, " ")
			mapName := removeWhitespace(splits[0])
			readFromMap := stringToMap[prevMapName]
			saveToMap := stringToMap[currentMapName]
			// fmt.Printf("About to swap prev: %s - %v for new map: %s - %v\n", prevMapName, readFromMap, currentMapName, saveToMap)
			// pass along ranges that didn't match
			// fmt.Printf("prev: %s, current: %s,  newmap: %s \n", prevMapName, currentMapName, mapName)
			if currentMapName != "initial-map" {
				for _, sr := range readFromMap {
					fmt.Println(sr)
					if !sr.hasBeenMapped {
						fmt.Printf("Adding %v to new map\n", sr)
						saveToMap = append(saveToMap, sr)
					}
				}
			}

			fmt.Printf("After filling in ranges that didnt match: %s - %v for new map: %s - %v\n", prevMapName, readFromMap, mapName, saveToMap)

			// save map
			stringToMap[currentMapName] = saveToMap
			// now swap the maps
			prevMapName = currentMapName
			currentMapName = mapName
		} else if len(removeWhitespace(line)) > 1 {
			ProcessLine(line, currentMapName, prevMapName, stringToMap)
		}
	}

	// some seeds might need to be added after final iteration
	readFromMap := stringToMap[prevMapName]
	saveToMap := stringToMap[currentMapName]

	if currentMapName != "initial-map" {
		for _, sr := range readFromMap {
			fmt.Println(sr)
			if !sr.hasBeenMapped {
				fmt.Printf("Adding %v to new map\n", sr)
				saveToMap = append(saveToMap, sr)
			}
		}
	}
	stringToMap[currentMapName] = saveToMap

	lowestLoc := 1000000000000000000
	fmt.Println("len humidity map: ", len(stringToMap["humidity-to-location"]))
	for _, v := range stringToMap["humidity-to-location"] {
		fmt.Println(v)
		// add weird hack to find smallest seed
		if v.start == 0 {
			continue
		}
		if v.start < lowestLoc {
			lowestLoc = v.start
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Lowest Location", lowestLoc)
}
