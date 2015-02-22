package main

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"reflect"
	//log "github.com/Sirupsen/logrus"
	//"regexp"
	"strings"
)

var OPERATORS = []string{"==", "!="}

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
			fmt.Println("parts:", parts)
			if len(parts) != 2 {
				return nil, fmt.Errorf("Value %s is not valid.", arg)
			}
			exprs = append(exprs, expr{key: parts[0], operator: op, value: parts[1]})
			found = true
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
