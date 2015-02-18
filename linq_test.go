package main

import (
	"bytes"
	linq "github.com/ahmetalpbalkan/go-linq"
	"html/template"
	"testing"
)

type Student struct {
	id, age int
	Name    string
}

func TestLINGInTemplate(t *testing.T) {
	var err error

	clara := &Student{id: 1, age: 19, Name: "clara"}
	tim := &Student{id: 1, age: 17, Name: "tim"}

	students := []*Student{clara, tim}

	Filter := func(queryOrT interface{}) linq.Query {
		whereFunc := func(s linq.T) (bool, error) {
			return s.(*Student).age >= 18, nil
		}

		var query linq.Query
		if _, ok := queryOrT.(linq.Query); ok {
			query = queryOrT.(linq.Query)
		} else {
			query = linq.From(queryOrT)
		}

		query = query.Where(whereFunc)
		return query
	}

	Sort := func(query linq.Query) linq.Query {
		return query
	}

	Results := func(query linq.Query) []linq.T {
		results, err := query.Results()
		if err != nil {
			panic(err)
		}
		return results

		/*
			var students []*Student
			for _, r := range results {
				students = append(students, r.(*Student))
			}

			return students
		*/
	}

	funcMap := template.FuncMap{
		"Filter":  Filter,
		"Sort":    Sort,
		"Results": Results,
	}
	tmpl := template.New("test")
	tmpl = tmpl.Funcs(funcMap)
	tmpl, err = tmpl.Parse(`
	{{ $results :=  . | Filter | Results }}
	Results:
	{{ range $index,$ele := $results }}
		{{ $index }} - {{ $ele.Name }} 
	{{ end }}
	End
	`)
	//tmpl, err = tmpl.Parse(`{{ From(.).Where(func (s T) (bool,error) { return s(*Student).age >= 18,nil }).Results() }}`)
	if err != nil {
		t.Error(err)
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, students)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", output.String())

}
