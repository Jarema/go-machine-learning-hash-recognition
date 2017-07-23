package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
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

	csvData, err := gocsv.MarshalBytes(dataset)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("extracted_data.csv", csvData, 0644)

	rawData, err := base.ParseCSVToInstances("extracted_data.csv", true)
	if err != nil {
		panic(err)
	}

	cls := knn.NewKnnClassifier("euclidean", "linear", 5)
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.70)

	cls.Fit(trainData)

	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))

	// checking with different example

	exampleInputFile, err := ioutil.ReadFile("example.csv")
	if err != nil {
		panic(err)
	}
	var eiData []inputRow
	if err := gocsv.Unmarshal(bytes.NewReader(exampleInputFile), &eiData); err != nil {
		panic(err)
	}
	var edataset []extractedRow
	for _, v := range eiData {
		row := extractFeatures(v.Value)
		row.Category = v.Type
		edataset = append(edataset, row)
	}

	ecsvData, err := gocsv.MarshalBytes(edataset)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("example_extracted_data.csv", ecsvData, 0644)
	if err != nil {
		panic(err)
	}

	exampleData, err := base.ParseCSVToTemplatedInstances("example_extracted_data.csv", true, rawData)
	if err != nil {
		panic(err)
	}
	check, err := cls.Predict(exampleData)
	if err != nil {
		panic(err)
	}

	exampleFile, err := ioutil.ReadFile("example.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(bytes.NewReader(exampleFile))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	_, size := check.Size()
	for i := 0; i < size; i++ {
		headers := records[0]
		fmt.Printf("%v:%v, type: %v\n", headers[0], records[i+1][0], check.RowString(i))
	}

}
