package model

import (
	"fmt"
	"time"
)

type Employee struct {
	Name     string
	LastName string
}

type Vacation struct {
	Employee
	Start time.Time
	End   time.Time
}

type EmployeeStat struct {
	Employee
	TotalDays int
}

// SHOULD BE CLARIFIED:
// IF END_DATE == START_DATE =>  INCORRECT DATA OR 1 DAY VACATION ??
func (vac *Vacation) VacationDaysCount() int {
	return int(vac.End.Sub(vac.Start).Hours() / 24) // +1
}

func (emp Employee) String() string {
	return fmt.Sprintf("%s %s", emp.Name, emp.LastName)
}

func (vac Vacation) String() string {
	return fmt.Sprintf("%s, %s %s", vac.Employee.String(), vac.Start.Format(time.RFC822), vac.End.Format(time.RFC822))
}

func (eStat EmployeeStat) String() string {
	return fmt.Sprintf("%s: %d", eStat.Employee.String(), eStat.TotalDays)
}
