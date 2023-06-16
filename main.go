package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/x12"
)

func readData(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: x12decode <filename>")
		os.Exit(1)
	}
	input_filename := argsWithoutProg[0]
	data, err := readData(input_filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	doc, err := x12.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	jsonOutput, err := json.MarshalIndent(doc, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output_filename := fileNameWithoutExtSliceNotation(input_filename) + ".json"
	err = ioutil.WriteFile(output_filename, jsonOutput, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output written to %s\n.", output_filename)
}
