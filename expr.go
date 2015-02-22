package main

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	//log "github.com/Sirupsen/logrus"
	//"regexp"
	"strings"
)

var OPERATORS = []string{"==", "!=", "in"}

type expr struct {
	key      string
	operator int
	value    string
}

func parseExprs(args []string) ([]expr, error) {
	exprs := []expr{}

	for _, arg := range args {
		found := false
		for i, op := range OPERATORS {
			//split with the op
			parts := strings.SplitN(arg, op, 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("Value %s is not valid.", arg)
			}
			exprs = append(exprs, expr{key: parts[0], operator: i, value: parts[1]})
			found = true
		}
		if !found {
			return nil, fmt.Errorf("One of operator ==,!= is expected")
		}
	}

	return exprs, nil
}

func (e expr) WhereFunc(files Files) func(linq.T) (bool, error) {
	whereFunc := func(linq.T) (bool, error) {
		return false, nil
	}
	return whereFunc
}
