package graph

import (
	"context"
	"reflect"
	"testing"

	"github.com/LuigiAzevedo/graphql_min_roman_numeral/cmd/graph/model"
)

func TestMutation_Search(t *testing.T) {
	app := mutationResolver{&Resolver{}}
	ctx := context.Background()

	tests := map[string]struct {
		input      string
		wantNumber string
		wantValue  int
	}{
		"one prime number":     {input: "AXIBIV", wantNumber: "XI", wantValue: 11},
		"two prime numbers":    {input: "XIOJKJKJIUVII", wantNumber: "VII", wantValue: 7},
		"decimal":              {input: "777", wantNumber: "", wantValue: 0},
		"big non prime number": {input: "M", wantNumber: "", wantValue: 0},
		"big prime number":     {input: "MCMLIHUAHDUHH", wantNumber: "MCMLI", wantValue: 1951},
		"empty":                {input: "", wantNumber: "", wantValue: 0},
		"non roman numbers":     {input: "ABFHKQZ", wantNumber: "", wantValue: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// call mutation function
			got, err := app.Search(ctx, model.Word{Text: test.input})
			if err != nil {
				t.Error(err)
			}

			// verify if input and result is the same fo number and value
			if !reflect.DeepEqual(test.wantNumber, *got.Number) {
				t.Fatalf("expected: %v, got: %v", test.wantNumber, *got.Number)
			}

			if !reflect.DeepEqual(test.wantValue, *got.Value) {
				t.Fatalf("expected: %v, got: %v", test.wantValue, *got.Value)
			}
		})
	}
}
