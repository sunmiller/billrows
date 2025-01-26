package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	file, err := os.Open("./data/measurements.txt")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer file.Close()
	// Convert from Windows-1252 (ISO-8859-1) to UTF-8
	decoder := charmap.Windows1252.NewDecoder()
	reader := transform.NewReader(file, decoder)

	fileOutputMap := make(map[string][]float32)

	scanner := bufio.NewScanner(reader)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		cityName, cityTempValue := processLine(line)
		// Check if the key exists
		if _, exists := fileOutputMap[cityName]; exists {
			// Convert string to float64
			floatValue, err := strconv.ParseFloat(cityTempValue, 32)
			if err != nil {
				fmt.Printf("Error converting string to float64: %v\n", err)
				return
			}
			fileOutputMap[cityName] = append(fileOutputMap[cityName], float32(floatValue))
		} else {
			// Key does not exist, add it with a default value
			// Convert string to float64
			floatValue, err := strconv.ParseFloat(cityTempValue, 32)
			if err != nil {
				fmt.Printf("Error converting string to float64: %v\n", err)
				return
			}
			fileOutputMap[cityName] = []float32{float32(floatValue)}

		}
	}
	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// printSortedList(&fileOutputMap)
}

func printSortedList(fileOutputMap *map[string][]float32) {
	keys := make([]string, 0, len(*fileOutputMap))
	for key := range *fileOutputMap {
		keys = append(keys, key)
	}
	sort.Strings(keys) // Sort keys alphabetically

	for _, key := range keys {
		// Print the key and its sorted values
		fmt.Println(key, (*fileOutputMap)[key])
	}
}

func processLine(line string) (string, string) {
	splitResult := strings.Split(line, ";")

	return splitResult[0], splitResult[1]
}
