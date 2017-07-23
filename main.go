package main

import (
	"bytes"
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

	csv, err := gocsv.MarshalBytes(dataset)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("extracted_data.csv", csv, 0644)

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

}
