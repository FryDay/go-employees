package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/FryDay/go-employees/models"
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

	var emps models.Employees
	_, err := db.Select(&emps, "select * From employees")
	checkError(err)

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
	var emp models.Employee
	var empID int
	var err error

	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if empID, err = strconv.Atoi(vars["employeeID"]); err != nil {
		panic(err)
	}

	err = db.SelectOne(&emp, "select employee_id, name, title from employees where employee_id = ?", empID)
	checkError(err)

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(emp); err != nil {
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
		w.WriteHeader(http.StatusBadRequest)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	err = db.Insert(&emp)
	checkError(err)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(emp); err != nil {
		panic(err)
	}
}

func employeeChange(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	var employeeID int
	var err error

	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if employeeID, err = strconv.Atoi(vars["employeeID"]); err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &emp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	emp.ID = int64(employeeID)
	count, err := db.Update(&emp)
	checkError(err)

	if count > 0 {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(emp); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{code: http.StatusNotFound, text: "Not found"}); err != nil {
			panic(err)
		}
	}
}

func employeeDelete(w http.ResponseWriter, r *http.Request) {
	var employeeID int
	var err error

	vars := mux.Vars(r)

	if employeeID, err = strconv.Atoi(vars["employeeID"]); err != nil {
		panic(err)
	}

	emp := models.Employee{ID: int64(employeeID)}
	count, err := db.Delete(&emp)

	if count > 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Employee deleted")
	} else {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err = json.NewEncoder(w).Encode(jsonErr{code: http.StatusNotFound, text: "Not found"}); err != nil {
			panic(err)
		}
	}
}
