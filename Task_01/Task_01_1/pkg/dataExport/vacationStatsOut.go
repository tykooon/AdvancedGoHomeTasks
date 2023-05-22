package dataExport

import (
	"fmt"
	"os"

	"example.com/tykoon/pkg/model"
)

type VacationStatsFileWriter struct {
	_outputFilePath string
}

func NewVacationStatsFileWriter(outputFile string) *VacationStatsFileWriter {
	return &VacationStatsFileWriter{_outputFilePath: outputFile}
}

func (w *VacationStatsFileWriter) WriteDurationStat(data model.EmployeeStatSlice) error {
	file, err := os.Create(w._outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, val := range data {
		_, err = fmt.Fprintf(file, "%s: %d\n", val.Employee.String(), val.TotalDays)
		if err != nil {
			break
		}
	}

	return err
}
