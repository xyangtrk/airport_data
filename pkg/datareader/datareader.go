package datareader

import (
	"encoding/json"
	"io"
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

