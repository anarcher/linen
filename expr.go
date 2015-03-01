package main

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"reflect"
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
		keyValue := reflect.ValueOf(file).FieldByName(e.key)

		//TODO: This converting check is not good. I will be digging about how do golang/reflect works.
		switch e.operator {
		case "==":
			if keyValue.String() == e.value {
				return true, nil
			}
		case "!=":
			if keyValue.String() != e.value {
				return true, nil
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
