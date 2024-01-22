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
				// fill in the map with ranges
				for i := 0; i < rangeLen; i++ {
					currentMap[sourceRangeStart+i] = destRangeStart + i
				}
				// update map
				stringToMap[currentMapName] = currentMap
			}
		}
	}
	lowestLoc := 100000
	// now that all the maps have been parsed, step through each map
	for _, s := range seeds {
		soil, ok := stringToMap["seed-to-soil"][s]
		if !ok {
			soil = s
		}
		fert, ok := stringToMap["soil-to-fertilizer"][soil]
		if !ok {
			fert = soil
		}
		water, ok := stringToMap["fertilizer-to-water"][fert]
		if !ok {
			water = fert
		}
		light, ok := stringToMap["water-to-light"][water]
		if !ok {
			light = water
		}
		temp, ok := stringToMap["light-to-temperature"][light]
		if !ok {
			temp = light
		}
		hum, ok := stringToMap["temperature-to-humidity"][temp]
		if !ok {
			hum = temp
		}
		loc, ok := stringToMap["humidity-to-location"][hum]
		if !ok {
			loc = hum
		}
		if loc < lowestLoc {
			lowestLoc = loc
		}
		fmt.Println("Found location: ", loc)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println("Lowest Location", lowestLoc)
}
