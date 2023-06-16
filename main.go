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

type Input struct {
	ReportContent string `json:"report_content"`
}

var sampleDoc = `ISA*00*          *00*          *08*9254110060     *ZZ*123456789      *041216*0805*U*00501*000095071*0*P*>~
GS*AG*5137624388*123456789*20041216*0805*95071*X*005010~
ST*824*021390001*005010X186A1~
BGN*11*FFA.ABCDEF.123456*20020709*0932**123456789**WQ~
N1*41*ABC INSURANCE*46*111111111~
PER*IC*JOHN JOHNSON*TE*8005551212*EX*1439~
N1*40*SMITHCO*46*A1234~
OTI*TA*TN*NA***20020709*0902*2*0001*834*005010X220A1~
SE*7*021390001~
GE*1*95071~
IEA*1*000095071~`

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
