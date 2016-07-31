package models

type Employee struct {
	ID    int64  `json:"id" db:"employee_id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Employees []Employee
