package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func OpenJsonFile(filename string) *os.File {
	jsonFile, err := os.Open(filename)
	// File Not Read Quiz Can't Continue
	if err != nil {
		log.Fatal("File Not Read")
	}

	return jsonFile
}

func ParseFile(filename string) Quiz {
	jsonFile := OpenJsonFile(filename)
	byteArray, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var quiz Quiz
	json.Unmarshal(byteArray, &quiz)
	defer jsonFile.Close()
	// Initializing all values to zero
	for i := range quiz.Quiz {
		score_map[i] = 0
	}
	return quiz
}
