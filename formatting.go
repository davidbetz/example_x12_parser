package main

import (
	"encoding/json"
	"encoding/xml"
)

type formatter interface {
	Format(v any) ([]byte, error)
	Name() string
}

type xmlFormatter struct {
}

func (t *xmlFormatter) Name() string {
	return "xml"
}

func (t *xmlFormatter) Format(input any) ([]byte, error) {
	output, err := xml.MarshalIndent(input, "", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

type jsonFormatter struct {
}

func (t *jsonFormatter) Name() string {
	return "json"
}

func (t *jsonFormatter) Format(input any) ([]byte, error) {
	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		return nil, err
	}
	return output, nil
}
