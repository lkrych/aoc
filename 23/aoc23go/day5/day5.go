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

// UpdateMapNames updates the current and previous map names.
func UpdateMapNames(currentMapName, prevMapName, mapName string) (string, string) {
	prevMapName = currentMapName
	currentMapName = mapName
	return currentMapName, prevMapName
}

// parseLine parses a line into three integers.
func parseLine(line string) (int, int, int) {
	parts := strings.Split(line, " ")
	destRangeStart := convertStringToInt(parts[0])
	sourceRangeStart := convertStringToInt(parts[1])
	rangeLen := convertStringToInt(parts[2])
	return destRangeStart, sourceRangeStart, rangeLen
}

// printSeedRangeInfo prints information about a SeedRange.
func printSeedRangeInfo(sr SeedRange, prevMapName string, sourceRangeStart, rangeLen int) {
	fmt.Printf("Reading %v from %s and comparing to start: %d, len: %d\n", sr, prevMapName, sourceRangeStart, rangeLen)
}

// handleSeedRange processes a SeedRange and returns an updated range and new ranges.
func handleSeedRange(sr SeedRange, mapStart, mapLen, destRangeStart int) (SeedRange, []SeedRange, []SeedRange) {
	var seedRangesToSave []SeedRange
	var seedRangesToRead []SeedRange

	// Implement the logic for handling the SeedRange here.
	// This involves checking the conditions as in the original code and updating
	// the SeedRange and creating new SeedRanges as necessary.

	//srFinish needs to subtract 1 because count includes sr.start
	srFinish := (sr.start + sr.len) - 1
	mapFinish := (mapStart + mapLen) - 1
	if sr.start >= mapStart && sr.start <= mapFinish && srFinish <= mapFinish {
		// fmt.Println("Assigning ", newDest, " to ", currentMapName, " for seed ", seed)
		diff := sr.start - mapStart
		newDest := destRangeStart + diff
		seedRangesToSave = append(seedRangesToSave, SeedRange{start: newDest, len: sr.len, hasBeenMapped: false})
		fmt.Printf("Case 1: overlaps entire range: %v\n ", SeedRange{start: newDest, len: sr.len})
		sr.hasBeenMapped = true
	} else if sr.start <= mapStart && srFinish >= mapStart && srFinish <= mapFinish {
		// case 2 sr starts before map range
		// this means we need to save the unmapped values up to the mapStart
		nonOverlapLen := mapStart - sr.start
		seedRangesToSave = append(seedRangesToSave, SeedRange{start: sr.start, len: nonOverlapLen, hasBeenMapped: false})
		fmt.Printf("Case 2: unmapped values: %v\n ", SeedRange{start: sr.start, len: nonOverlapLen, hasBeenMapped: false})
		// push unmapped values back into iterator list
		seedRangesToRead = append(seedRangesToRead, SeedRange{start: sr.start, len: nonOverlapLen, hasBeenMapped: false})

		// we also need to save the overlap from mapStart
		overLapLen := (srFinish) - mapStart + 1
		seedRangesToSave = append(seedRangesToSave, SeedRange{start: destRangeStart, len: overLapLen, hasBeenMapped: false})
		fmt.Printf("Case 2: overlapped values: %v\n ", SeedRange{start: destRangeStart, len: overLapLen})
		sr.hasBeenMapped = true
	} else if sr.start >= mapStart && sr.start <= mapFinish && srFinish >= mapFinish {
		// case 3 sr starts after map range and overlaps over the range
		// this means we need to save the unmapped values after the sourceRange finish
		nonOverlapStart := mapStart + mapLen
		nonOverlapLen := srFinish - nonOverlapStart + 1
		seedRangesToSave = append(seedRangesToSave, SeedRange{start: nonOverlapStart, len: nonOverlapLen, hasBeenMapped: false})
		fmt.Printf("Case 3: unmapped values: %v\n ", SeedRange{start: nonOverlapStart, len: nonOverlapLen})
		// push unmapped values back into iterator list
		seedRangesToRead = append(seedRangesToRead, SeedRange{start: nonOverlapStart, len: nonOverlapLen, hasBeenMapped: false})

		// we also need to save the overlap from mapStart
		diff := sr.start - mapStart
		newDest := destRangeStart + diff
		newLen := (mapStart + mapLen) - sr.start
		seedRangesToSave = append(seedRangesToSave, SeedRange{start: newDest, len: newLen})
		fmt.Printf("Case 3: overlapped values: %v\n ", SeedRange{start: newDest, len: newLen, hasBeenMapped: false})
		sr.hasBeenMapped = true
	}

	return sr, seedRangesToSave, seedRangesToRead
}

// ProcessLine processes each line and updates the maps accordingly.
func ProcessLine(line string, currentMapName, prevMapName string, stringToMap map[string][]SeedRange) {
	if len(removeWhitespace(line)) <= 1 {
		return
	}

	readFromMap, saveToMap := stringToMap[prevMapName], stringToMap[currentMapName]
	destRangeStart, sourceRangeStart, rangeLen := parseLine(line)

	addToReadFromMap := []SeedRange{}
	for idx, seedRange := range readFromMap {
		if seedRange.hasBeenMapped {
			continue
		}

		printSeedRangeInfo(seedRange, prevMapName, sourceRangeStart, rangeLen)
		updatedSeedRange, toSaveRanges, toAppendRanges := handleSeedRange(seedRange, sourceRangeStart, rangeLen, destRangeStart)

		if updatedSeedRange.hasBeenMapped {
			readFromMap[idx] = updatedSeedRange
			saveToMap = append(saveToMap, toSaveRanges...)
		}

		if len(toAppendRanges) > 0 {
			addToReadFromMap = append(addToReadFromMap, toAppendRanges...)
		}
	}

	readFromMap = append(readFromMap, addToReadFromMap...)
	stringToMap[currentMapName], stringToMap[prevMapName] = saveToMap, readFromMap
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
			stringToMap["initial-map"] = initialMap
		} else if strings.Contains(line, "map") {
			splits := strings.Split(line, " ")
			mapName := removeWhitespace(splits[0])
			readFromMap := stringToMap[prevMapName]
			saveToMap := stringToMap[currentMapName]
			// pass along ranges that didn't match
			for _, sr := range readFromMap {
				if !sr.hasBeenMapped {
					saveToMap = append(saveToMap, sr)
				}
			}
			// save map
			stringToMap[currentMapName] = saveToMap
			currentMapName, prevMapName = UpdateMapNames(currentMapName, prevMapName, mapName)
		} else if len(removeWhitespace(line)) > 1 {
			ProcessLine(line, currentMapName, prevMapName, stringToMap)
		}
	}

	lowestLoc := 1000000000000000000
	for _, v := range stringToMap["humidity-to-location"] {
		fmt.Println(v)
		if v.start < lowestLoc {
			lowestLoc = v.start
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Lowest Location", lowestLoc)
}
