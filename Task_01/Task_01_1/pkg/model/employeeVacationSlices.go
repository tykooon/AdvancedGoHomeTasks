package model

type VacationSlice []Vacation

type EmployeeStatSlice []EmployeeStat

func (emps EmployeeStatSlice) Len() int {
	return len(emps)
}

func (emps EmployeeStatSlice) Less(i, j int) bool {
	if emps[i].TotalDays < emps[j].TotalDays {
		return false
	} else if emps[i].TotalDays > emps[j].TotalDays {
		return true
	}
	return emps[i].LastName < emps[j].LastName
}

func (emps EmployeeStatSlice) Swap(i, j int) {
	emps[i], emps[j] = emps[j], emps[i]
}

func MakeStatsFromVacations(all VacationSlice) EmployeeStatSlice {
	dict := make(map[Employee]int, 0)
	for _, val := range all {
		dict[val.Employee] += val.VacationDaysCount()
	}
	res := make([]EmployeeStat, 0)
	for k, v := range dict {
		res = append(res, EmployeeStat{Employee: k, TotalDays: v})
	}
	return res
}
