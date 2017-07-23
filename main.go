package main

import (
	"bytes"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
)

type inputRow struct {
	Value string `csv:"value"`
	Type  string `csv:"type"`
}

func main() {

	inputFile, err := ioutil.ReadFile("input.csv")
	if err != nil {
		panic(err)
	}
	var iData []inputRow
	if err := gocsv.Unmarshal(bytes.NewReader(inputFile), &iData); err != nil {
		panic(err)
	}

	var dataset []extractedRow
	for _, v := range iData {
		row := extractFeatures(v.Value)
		row.Category = v.Type
		dataset = append(dataset, row)
	}
	fmt.Println(dataset)

	csv, err := gocsv.MarshalBytes(dataset)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("extracted_data.csv", csv, 0644)
}
