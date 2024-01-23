package day5

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

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

// create function that can be done in individual thread
func processSeed(filepath string, seedSrc int, seedRange int) int {
	scanner, err := input.ReadInputFile(filepath)
	if err != nil {
		panic(err)
	}
	defer scanner.Scan() // Close the file when done reading

	// BEGIN CODING FOR DAY HERE
	// INITIALIZE GLOBAL VALUES
	seeds := []int{}
	for j := seedSrc; j <= seedSrc+seedRange; j++ {
		seeds = append(seeds, j)
	}
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

	currentMapName := "seed-to-soil"
	hasBeenMapped := map[int]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		// load the current working map
		if strings.Contains(line, "seeds") {
			continue
		} else if strings.Contains(line, "map") {
			splits := strings.Split(line, " ")
			// fmt.Println(splits)
			mapName := removeWhitespace(splits[0])
			// fmt.Println("Updating map: ", mapName)
			// seed next map with previous map
			for _, s := range seeds {
				prevMap := stringToMap[currentMapName]
				prevVal, ok := prevMap[s]
				if !ok {
					// if a previous value wasn't found default to seed value
					prevVal = s
					// fmt.Println("Defaulting to seed value for map ", mapName)
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
				// fmt.Println("Parsing ", currentMapName, " dest: ", destRangeStart, ", src: ", sourceRangeStart, ", range: ", rangeLen)
				// iterate through each seed value in currentMap
				for seed, v := range currentMap {
					hasAlreadyBeenMapped := hasBeenMapped[seed]
					if v >= sourceRangeStart && v <= sourceRangeStart+rangeLen && !hasAlreadyBeenMapped {
						// find the difference
						diff := v - sourceRangeStart
						newDest := destRangeStart + diff
						// fmt.Println("Assigning ", newDest, " to ", currentMapName, " for seed ", seed)
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
	return lowestLoc
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

	type ResultType struct {
		Value int // Replace with the actual type of your result
	}

	// Shared data structure to store results
	results := make(map[int]ResultType) // Key is the seed
	mutex := &sync.Mutex{}
	var wg sync.WaitGroup

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
				wg.Add(1)
				go func(fp string, seedS int, seedR int, i int) {
					defer wg.Done()
					// Perform computation for seed 's' and obtain result
					result := ResultType{Value: processSeed(fp, seedS, seedR)}

					// Store the result in the shared map
					mutex.Lock()
					results[i] = result
					mutex.Unlock()
				}(filepath, seedSrc, seedRange, i)
			}
			break
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	lowestLoc := 1000000000000000000
	for _, v := range results {
		if v.Value < lowestLoc {
			lowestLoc = v.Value
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Lowest Location", lowestLoc)
}
