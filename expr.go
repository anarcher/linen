package main

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"strings"
	//log "github.com/Sirupsen/logrus"
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
				exprs = append(exprs, expr{key: strings.TrimSpace(parts[0]), operator: op, value: parts[1]})
				found = true
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
		if !keyValue.IsValid() {
			return false, nil
		}

		op, err := NewOperator(keyValue.Interface(), e.value)
		if err != nil {
			return false, err
		}

		switch e.operator {
		case "==":
			if op.Eq() == true {
				return true, nil
			}
		case "!=":
			if op.Eq() == false {
				return true, nil
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
		if !thisVal.IsValid() || !thisVal.IsValid() {
			return false
		}

		op, err := NewOperator(thisVal.Interface(), thatVal.Interface())
		if err != nil {
			fmt.Println(err)
			return false
		}

		opRes := op.Cmp()

		//TODO: It's really not Good.
		switch e.operator {
		case "asc":
			if opRes == -1 {
				return true
			}
			return false
		case "desc":
			if opRes == 1 {
				return true
			}
			return false
		}
		return false
	}
	return orderByFunc
}
