package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/FryDay/employees/models"
	"github.com/gorilla/mux"
)

type jsonErr struct {
	code int
	text string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func employeeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	rows, err := db.Query("select * From employees")
	checkError(err)

	var emps models.Employees
	for rows.Next() {
		var emp models.Employee
		checkError(rows.Scan(&emp.ID, &emp.Name, &emp.Title))
		emps = append(emps, emp)
	}

	if len(emps) > 0 {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(emps); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{code: http.StatusNotFound, text: "No employees"}); err != nil {
		panic(err)
	}
}

func employeeShow(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	var employeeID int
	var err error

	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if employeeID, err = strconv.Atoi(vars["employeeID"]); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("select id, name, title from employees where id = ?")
	checkError(err)
	defer stmt.Close()

	if err = stmt.QueryRow(employeeID).Scan(&employee.ID, &employee.Name, &employee.Title); err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err = json.NewEncoder(w).Encode(jsonErr{code: http.StatusNotFound, text: "No employees"}); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(employee); err != nil {
		panic(err)
	}
}

func employeeCreate(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &emp); err != nil {
		w.WriteHeader(422) // unprocessable
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	stmt, err := db.Prepare("insert into employees(name, title) values (?,?)")
	checkError(err)
	defer stmt.Close()

	result, err := stmt.Exec(emp.Name, emp.Title)
	checkError(err)

	id, err := result.LastInsertId()
	checkError(err)

	emp.ID = id
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(emp); err != nil {
		panic(err)
	}
}