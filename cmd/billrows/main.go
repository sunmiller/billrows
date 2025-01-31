package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cityTemps struct {
	min, max, sum float64
	count         int64
}

func main() {

	file, err := os.Open("./data/measurements.txt")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer file.Close()

	fileOutputMap := make(map[string]*cityTemps)

	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// cityName, cityTempValue := splitLine(line)
		cityName, cityTemp, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}
		cityTempValue, err := strconv.ParseFloat(cityTemp, 64)
		if err != nil {
			panic(err)
		}
		// Check if the key exists
		foundCityTemp, ok := fileOutputMap[cityName]
		if ok {
			foundCityTemp.min = min(foundCityTemp.min, cityTempValue)
			foundCityTemp.max = max(foundCityTemp.max, cityTempValue)
			foundCityTemp.sum += cityTempValue
			foundCityTemp.count++
		} else {
			fileOutputMap[cityName] = &cityTemps{
				min:   cityTempValue,
				max:   cityTempValue,
				sum:   cityTempValue,
				count: 1,
			}

		}

	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	printSortedList(&fileOutputMap)
}

func printSortedList(fileOutputMap *map[string]*cityTemps) {
	cities := make([]string, 0, len(*fileOutputMap))
	for city := range *fileOutputMap {
		cities = append(cities, city)
	}
	sort.Strings(cities) // Sort keys alphabetically

	for _, key := range cities {
		// Print the key and its sorted values

		mean := (*fileOutputMap)[key].sum / float64((*fileOutputMap)[key].count)

		fmt.Printf("%s=%.1f/%.1f/%.1f,", key, (*fileOutputMap)[key].min, mean, (*fileOutputMap)[key].max)

	}
}
