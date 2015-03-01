package main

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"reflect"
	"strconv"
	//log "github.com/Sirupsen/logrus"
	//"regexp"
	"strings"
)

var OPERATORS = []string{"==", "!=", "asc", "desc"}

type expr struct {
	key      string
	operator string
	value    string
}

func parseExprs(args []string) ([]expr, error) {
	exprs := []expr{}

	for _, arg := range args {
		found := false
		for _, op := range OPERATORS {
			//split with the op
			parts := strings.SplitN(arg, op, 2)
			if len(parts) == 2 {
				exprs = append(exprs, expr{key: parts[0], operator: op, value: parts[1]})
				found = true
			} else if len(parts) == 1 {
				exprs = append(exprs, expr{key: parts[0], operator: op})
				found = true
			}
			if found {
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("One of operator ==,!= is expected")
		}
	}

	return exprs, nil
}

func (e expr) WhereFunc() func(linq.T) (bool, error) {

	whereFunc := func(t linq.T) (bool, error) {
		file := *(t.(*File))
		var keyValue reflect.Value
		if strings.HasPrefix(e.key, "Meta.") {
			key := strings.Replace(e.key, "Meta.", "", 1)
			if val, ok := file.Meta[key]; ok == true {
				keyValue = reflect.ValueOf(val)
			} else {
				return false, nil
			}

		} else {
			keyValue = reflect.ValueOf(file).FieldByName(e.key)
		}

		switch e.operator {
		case "==":
			switch keyValue.Kind() {
			case reflect.Int:
				eVal, err := strconv.Atoi(e.value)
				if err != nil {
					return false, err
				}
				if keyValue.Interface().(int) == eVal {
					return true, nil
				}
			case reflect.Bool:
				eVal, err := strconv.ParseBool(e.value)
				if err != nil {
					return false, err
				}
				if keyValue.Interface().(bool) == eVal {
					return true, nil
				}
			case reflect.String:
				if keyValue.Interface().(string) == e.value {
					return true, nil
				}
			}

		case "!=":
			switch keyValue.Kind() {
			case reflect.Int:
				eVal, err := strconv.Atoi(e.value)
				if err != nil {
					return false, err
				}
				if keyValue.Interface().(int) == eVal {
					return true, nil
				}
			case reflect.Bool:
				eVal, err := strconv.ParseBool(e.value)
				if err != nil {
					return false, err
				}
				if keyValue.Interface().(bool) != eVal {
					return true, nil
				}
			case reflect.String:
				if keyValue.Interface().(string) != e.value {
					return true, nil
				}
			}
		}

		return false, nil
	}
	return whereFunc
}

func (e expr) OrderByFunc() func(this linq.T, that linq.T) bool {
	orderByFunc := func(this linq.T, that linq.T) bool {
		thisF := *(this.(*File))
		thatF := *(that.(*File))
		thisVal := reflect.ValueOf(thisF).FieldByName(e.key)
		thatVal := reflect.ValueOf(thatF).FieldByName(e.key)

		//TODO: It's really not Good.
		switch e.operator {
		case "asc":
			return thisVal.String() > thatVal.String()
		case "desc":
			return thatVal.String() < thatVal.String()
		}
		return false
	}
	return orderByFunc
}
