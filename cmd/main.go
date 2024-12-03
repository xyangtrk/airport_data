package main

import (
	"log"

	"github.com/xyangtrk/airport_data/pkg/datareader"
	// "github.com/xyangtrk/airport_data/pkg/redisclient"
)

func main() {
	// inputFilePath := "data/originaldata.json"
	// outputFilePath := "data/outputdata.json"

	// outputFilePath := "data/outputdata-test.json"

	// do not need to run this if we have the outputdata.json already
	// err := jsonparser.ParseAndBuildOutputData(inputFilePath, outputFilePath)
	//
	// if err != nil {
	//   log.Fatalf("Error in main: %v", err)
	// }

	// outputMap, err := datareader.ReadOutputData(outputFilePath)
	// if err != nil {
	// 	log.Fatalf("Error in main while reading output data: %v", err)
	// }

	// // Initialize Redis client
	// rdb := redisclient.InitializeRedisClient()

	// err = redisclient.InsertIntoRedis(outputMap, rdb)
	// if err != nil {
	// 	log.Fatalf("Error in main while inserting data into Redis: %v", err)
	// }

	// airline codes.
	inputFilePath := "data/original-airlines-code.json"
	outputFilePath := "data/output-airlines-code.json"

	if err := datareader.ParseAndBuildOutputAirlinesCodeData(inputFilePath, outputFilePath); err != nil {
		log.Fatalf("Error processing airlines code data: %v", err)
	}
}
