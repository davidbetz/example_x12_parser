package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/x12"
)

var (
	baseOutputPath = "./output"
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

func writeAll(folder string, formatter formatter, doc *x12.X12Document) error {
	output_filename := filepath.Join(folder, "output."+formatter.Name())
	fmt.Println(output_filename)
	output, err := formatter.Format(doc)
	if err != nil {
		return err
	}
	err = os.WriteFile(output_filename, output, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Output written to %s\n.", output_filename)
	return nil
}

func writeEach(folder string, formatter formatter, fgs []*x12.FunctionGroup) error {
	fmt.Printf("len(doc.Interchange.FunctionGroups): %d\n", len(fgs))
	for _, fg := range fgs {
		fmt.Printf("len(fg.Transactions): %d\n", len(fg.Transactions))
		for _, t := range fg.Transactions {
			output, err := formatter.Format(t)
			if err != nil {
				return err
			}
			output_filename := filepath.Join(folder, fmt.Sprintf("%s.%s", t.Header.TransactionSetControlNumber, formatter.Name()))
			err = os.WriteFile(output_filename, output, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		errorf("Usage: x12decode <filename> [format]")
	}
	input_filename := argsWithoutProg[0]
	formatSetting := "xml"
	if len(argsWithoutProg) > 1 {
		formatSetting = strings.ToLower(argsWithoutProg[1])
	}
	fmt.Println(input_filename)
	fmt.Println(formatSetting)
	var formatter formatter
	if formatSetting == "xml" {
		formatter = &xmlFormatter{}
	} else {
		formatter = &jsonFormatter{}
	}
	data, err := readData(input_filename)
	if err != nil {
		errorf(err.Error())
	}
	doc, err := x12.Decode(bytes.NewReader(data))
	if err != nil {
		errorf(err.Error())
	}
	baseFilename := fileNameWithoutExtSliceNotation(input_filename)
	folder := filepath.Join(baseOutputPath, baseFilename)
	err = os.MkdirAll(folder, 0755)
	if err != nil {
		errorf(err.Error())
	}
	err = writeAll(folder, formatter, doc)
	if err != nil {
		errorf(err.Error())
	}
	err = writeEach(folder, formatter, doc.Interchange.FunctionGroups)
	if err != nil {
		errorf(err.Error())
	}
	original_filename := filepath.Join(folder, input_filename)
	err = os.WriteFile(original_filename, data, 0644)
	if err != nil {
		errorf(err.Error())
	}
}
