package main

import (
  "log"

  "github.com/xyangtrk/airport_data/pkg/jsonparser"
 )


func main() {
  inputFilePath := "data/originaldata.json"
  outputFilePath := "data/outputdata.json"


  err := jsonparser.ParseAndBuildOutputData(inputFilePath, outputFilePath)

  if err != nil {
    log.Fatalf("Error in main: %v", err)
  }
}
