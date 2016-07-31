package main

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []route{
	route{
		"Index",
		"GET",
		"/",
		index,
	},
	route{
		"EmployeeIndex",
		"GET",
		"/employees",
		employeeIndex,
	},
	route{
		"EmployeeShow",
		"GET",
		"/employees/{employeeID}",
		employeeShow,
	},
	route{
		"EmployeeCreate",
		"POST",
		"/employees",
		employeeCreate,
	},
	route{
		"EmployeeChange",
		"PUT",
		"/employees/{employeeID}",
		employeeChange,
	},
	route{
		"EmployeeDelete",
		"DELETE",
		"/employees/{employeeID}",
		employeeDelete,
	},
}
