package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/tykoon/pkg/dataProvider"
	"example.com/tykoon/pkg/knnModel"
)

const SourceFile string = "data.txt"
const AttributesCount int = 2
const NeighborsCount int = 3

func main() {

	dataP, err := dataProvider.NewDataFileReader(SourceFile, AttributesCount)
	if err != nil {
		fmt.Printf("Unable to access file %s\n", SourceFile)
		return
	}

	data, err := dataP.ReadAllEntities()
	if err != nil {
		fmt.Printf("Unable to get data from the file %s\n%s", SourceFile, err.Error())
		return
	}

	model := knnModel.NewKnnModel(
		knnModel.EuclideanDist,
		AttributesCount,
		NeighborsCount,
	)

	model.AddEntities(data...)

	input := make([]float64, AttributesCount)
	var temp float64

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input ", AttributesCount, " attribute values as float64. One value per line:")
		for k := range input {
			for {
				fmt.Printf("attr. %d : ", k+1)
				_, err := fmt.Scanf("%f", &temp)
				stdin.ReadString('\n')
				if err == nil {
					break
				}
				fmt.Println("Wrong input. Try again!")
			}
			input[k] = temp
		}
		className, err := model.Classify(input...)
		if err == nil {
			fmt.Printf("Sucsessfully classified as *** %s ***\n\n", className)
		} else {
			fmt.Println(className, err.Error())
		}
	}
}
