package jsonparser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xyangtrk/airport_data/pkg/models"
)

func ParseAndBuildOutputData(inputPath string, outputPath string) error {
	jsonFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("cannot open JSON file: %v", err)
	}
	defer jsonFile.Close()

	// outputMap := make(map[string]models.OutputAirport)
	//
	// decoder := json.NewDecoder(jsonFile)
	// if _, err := decoder.Token(); err != nil {
	// 	return fmt.Errorf("failed to read start of JSON: %v", err)
	// }

  byteValue, err := io.ReadAll(jsonFile)
    if err != nil {
        return fmt.Errorf("failed to read JSON file: %v", err)
    }

    var airportsData models.AirportsData
    if err := json.Unmarshal(byteValue, &airportsData); err != nil {
        return fmt.Errorf("failed to parse JSON: %v", err)
    }

	// for decoder.More() {
	// 	var airport models.Airport
	// 	if err := decoder.Decode(&airport); err != nil {
	// 		return fmt.Errorf("failed to decode JSON: %v", err)
	// 	}

		// outputMap[airport.FS] = models.OutputAirport{
		// 	UTCOffset:          airport.UtcOffsetHours,
		// 	Latitude:           airport.Latitude,
		// 	Longitude:          airport.Longitude,
		// 	TimeZoneRegionName: airport.TimeZoneRegionName,
		// }
	// }


  // Build the output map
    outputMap := make(map[string]models.OutputAirport)
    for _, airport := range airportsData.Airports {
        outputMap[airport.FS] = models.OutputAirport{
            UTCOffset:          airport.UtcOffsetHours,
            Latitude:           airport.Latitude,
            Longitude:          airport.Longitude,
            TimeZoneRegionName: airport.TimeZoneRegionName,
        }
    }

  // if _, err := decoder.Token(); err != nil {
  //   return fmt.Errorf("failed to read end of JSON: %v", err)
  // }

  outputJSON, err := json.MarshalIndent(outputMap, "", "  ")

  if err != nil {
    return fmt.Errorf("failed to marshal output JSON: %v", err)
  }

  err = os.WriteFile(outputPath, outputJSON, 0644)

  if err != nil {
    return fmt.Errorf("failed to write output JSON: %v", err)
  }

	log.Printf("Output successfully written to %s", outputPath)
	return nil
}
