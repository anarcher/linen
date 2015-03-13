package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Operator interface {
	Eq() bool
	Cmp() int
}

type IntOp struct {
	a, b int
}

type StringOp struct {
	a, b string
}

type BoolOp struct {
	a, b bool
}

type StringListOp struct {
	a []string
	b string
}

func (op *IntOp) Eq() bool {
	if op.a == op.b {
		return true
	}
	return false
}

func (op *IntOp) Cmp() int {
	if op.a > op.b {
		return 1
	} else if op.a == op.b {
		return 0
	}
	return -1
}

func (op *StringOp) Eq() bool {
	if op.a == op.b {
		return true
	}
	return false
}

func (op *StringOp) Cmp() int {
	if op.a > op.b {
		return 1
	} else if op.a == op.b {
		return 0
	}
	return -1
}

func (op *BoolOp) Eq() bool {
	if op.a == op.b {
		return true
	}
	return false
}

func (op *BoolOp) Cmp() int {
	if op.a == true && op.b != false {
		return 1
	} else if op.a == true && op.b == true {
		return 0
	}
	return -1
}

func (op *StringListOp) Eq() bool {
	for _, a := range op.a {
		if a == op.b {
			return true
		}
	}
	return false
}

func (op *StringListOp) Cmp() int {
	for _, a := range op.a {
		if a == op.b {
			return 1
		}
	}
	return 0
}

func NewOperator(A interface{}, B interface{}) (op Operator, err error) {
	switch A.(type) {
	case int:
		var b int
		if _, ok := B.(string); ok {
			b, err = strconv.Atoi(B.(string))
			if err != nil {
				return
			}
		}
		op = &IntOp{a: A.(int), b: b}
	case string:
		op = &StringOp{a: A.(string), b: B.(string)}
	case bool:
		var b bool
		if _, ok := B.(string); ok {
			b, err = strconv.ParseBool(B.(string))
			if err != nil {
				return
			}
		}

		op = &BoolOp{a: A.(bool), b: b}
	case []interface{}:
		var a []string
		for _, i := range A.([]interface{}) {
			a = append(a, i.(string))
		}
		op = &StringListOp{a: a, b: B.(string)}
	default:
		err = fmt.Errorf("Can't match to supported type. %v %v", A, B)
	}

	return
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
