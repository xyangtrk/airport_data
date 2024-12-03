package datareader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/xyangtrk/airport_data/pkg/models"
)

func ReadOutputData(filePath string) (map[string]models.OutputAirport, error) {
	// Open the JSON file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// Read file contents
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// Parse the JSON into the output map
	var outputMap map[string]models.OutputAirport
	err = json.Unmarshal(byteValue, &outputMap)
	if err != nil {
		return nil, err
	}
	return outputMap, nil
}

// for airlines code.
func ParseAndBuildOutputAirlinesCodeData(inputFilePath, outputFilePath string) error {
	// Read the input file
	data, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	// Parse the JSON array
	var airlines []models.Airline
	if err := json.Unmarshal(data, &airlines); err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Create output map for filtered airlines
	outputMap := make(map[string]string)

	// Filter airlines and build output map
	for _, airline := range airlines {
		if airline.Active == "Y" && airline.Country == "United States" {
			if airline.IATA != "" && airline.ICAO != "" {
				outputMap[airline.IATA] = airline.ICAO
			}
		}
	}

	// Convert output map to JSON
	outputJSON, err := json.MarshalIndent(outputMap, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling output JSON: %v", err)
	}

	// Write to output file
	if err := ioutil.WriteFile(outputFilePath, outputJSON, 0644); err != nil {
		return fmt.Errorf("error writing output file: %v", err)
	}

	log.Printf("Successfully processed %d airlines and wrote output to %s", len(outputMap), outputFilePath)
	return nil
}
