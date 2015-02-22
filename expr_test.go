package main

import (
	"github.com/ahmetalpbalkan/go-linq"
	"testing"
)

func TestExpr(t *testing.T) {
	args := []string{"Path==/example"}
	exprs, err := parseExprs(args)
	if err != nil {
		t.Error(err)
	}

	t.Log(exprs)

	file := &File{Path: "/example", Ext: "", Meta: make(map[string]interface{})}
	files := []*File{file}
	query := linq.From(files)

	whereFunc := exprs[0].WhereFunc()
	query = query.Where(whereFunc)

	results, _err := query.Results()
	if _err != nil {
		t.Error(_err)
	}

	if len(results) != 1 {
		t.Error("result should has 1 file")
	}

	t.Log("results:", results)

}
