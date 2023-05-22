package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"example.com/tykoon/pkg/dataExport"
	"example.com/tykoon/pkg/dataProvider"
	"example.com/tykoon/pkg/model"
)

const SourceFileName string = "data.csv"
const OutputFileName string = "out.txt"
const ErrorOutputFileName string = "errors.txt"
const DateFormat string = "1/2/2006"

func main() {

	var errorWriter io.StringWriter

	errorWriter, err := dataExport.NewErrorFileWriter(ErrorOutputFileName)
	if err != nil {
		errorWriter = os.Stdout
	}

	var dataP = dataProvider.NewVacationsCsvReader(
		SourceFileName,
		DateFormat,
		errorWriter,
	)

	allVacations, err := dataP.GetAllVacations()
	if err != nil {
		fmt.Printf("Problems occured while getting data\n%s\n", err.Error())
		return
	}

	var vacationStats = model.MakeStatsFromVacations(allVacations)
	sort.Stable(vacationStats)

	fmt.Println("Output of vacation stats is started...")
	var output = dataExport.NewVacationStatsFileWriter(OutputFileName)

	err = output.WriteDurationStat(vacationStats)
	if err != nil {
		fmt.Println("Output stopped due to error: ", err.Error())
		return
	} else {
		fmt.Println("Output finished successfully.")
	}
}
