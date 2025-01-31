package fileone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./data/measurements.txt")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer file.Close()

	fileOutputMap := make(map[string][]float64)

	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// cityName, cityTempValue := splitLine(line)
		cityName, cityTempValue, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		// Check if the key exists
		if _, exists := fileOutputMap[cityName]; exists {
			// Convert string to float64
			floatValue, err := strconv.ParseFloat(cityTempValue, 64)
			if err != nil {
				fmt.Printf("Error converting string to float64: %v\n", err)
				return
			}
			fileOutputMap[cityName] = append(fileOutputMap[cityName], floatValue)
		} else {
			// Key does not exist, add it with a default value
			// Convert string to float64
			floatValue, err := strconv.ParseFloat(cityTempValue, 32)
			if err != nil {
				fmt.Printf("Error converting string to float64: %v\n", err)
				return
			}
			fileOutputMap[cityName] = []float64{floatValue}

		}
	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	printSortedList(&fileOutputMap)
}

func printSortedList(fileOutputMap *map[string][]float64) {
	keys := make([]string, 0, len(*fileOutputMap))
	for key := range *fileOutputMap {
		keys = append(keys, key)
	}
	sort.Strings(keys) // Sort keys alphabetically

	for _, key := range keys {
		// Print the key and its sorted values
		printCityWithCalculations(key, (*fileOutputMap)[key])
	}
}

func printCityWithCalculations(cityName string, temperatureSlices []float64) {
	// and the result values per station in the format <min>/<mean>/<max>, rounded to one fractional digit).
	var sum float64
	for _, value := range temperatureSlices {
		sum += value
	}
	mean := sum / float64(len(temperatureSlices))

	fmt.Printf("%s=%.1f/%.1f/%.1f,", cityName, slices.Min(temperatureSlices), mean, slices.Max(temperatureSlices))
}

// Split the line on ; and return the values.
func splitLine(line string) (string, string) {
	splitResult := strings.Split(line, ";")

	return splitResult[0], splitResult[1]
}
