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
		file := t.(*File)
		keyValue := getFileKeyValue(file, e.key)
		if keyValue == nil {
			return false, nil
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
		thisF := this.(*File)
		thatF := that.(*File)
		thisVal := getFileKeyValue(thisF, e.key)
		thatVal := getFileKeyValue(thatF, e.key)

		//TODO: It's really not Good.
		switch e.operator {
		case "asc":
			switch thisVal.Kind() {
			case reflect.Int:
				return thisVal.Interface().(int) > thatVal.Interface().(int)
			case reflect.Bool:
				if thisVal.Interface().(bool) == true {
					return true
				}
				return false
			case reflect.String:
				return thisVal.Interface().(string) > thatVal.Interface().(string)
			}
		case "desc":
			switch thisVal.Kind() {
			case reflect.Int:
				return thisVal.Interface().(int) < thatVal.Interface().(int)
			case reflect.Bool:
				if thisVal.Interface().(bool) == false {
					return true
				}
				return false
			case reflect.String:
				return thisVal.Interface().(string) < thatVal.Interface().(string)
			}
		}
		return false
	}
	return orderByFunc
}

func getFileKeyValue(file *File, key string) *reflect.Value {
	var keyValue reflect.Value
	if strings.HasPrefix(key, "Meta.") {
		key = strings.Replace(key, "Meta.", "", 1)
		if val, ok := file.Meta[key]; ok == true {
			keyValue = reflect.ValueOf(val)
		} else {
			return nil
		}
	} else {
		keyValue = reflect.ValueOf(*file).FieldByName(key)
	}
	return &keyValue
}
