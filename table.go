package testy

import (
	"reflect"
	"testing"
)

// Generator should return a test.
type Generator func(*testing.T) interface{}

// Table is a table of one or more tests to run against a single test function.
type Table map[string]Generator

// Add adds a single named test to the table.
func (tb Table) Add(name string, gen Generator) {
	if _, ok := tb[name]; ok {
		panic("Add(): Test " + name + " already defined.")
	}
	tb[name] = gen
}

// Run cycles through the defined tests, passing them one at a time to testFn.
// testFn must be a function which takes two arguments: *testing.T, and an
// arbitrary type, which must match the return value of the Generator functions.
func (tb Table) Run(t *testing.T, testFn interface{}) {
	testFnT := reflect.TypeOf(testFn)
	if testFnT.Kind() != reflect.Func {
		panic("testFn must be a function")
	}
	if testFnT.NumIn() != 2 || testFnT.In(0) != reflect.TypeOf(&testing.T{}) {
		panic("testFn must be of the form func(*testing.T, **)")
	}
	testType := reflect.TypeOf(testFn).In(1)
	testFnV := reflect.ValueOf(testFn)
	for name, genFn := range tb {
		t.Run(name, func(t *testing.T) {
			test := genFn(t)
			if reflect.TypeOf(test) != testType {
				t.Fatalf("Test generator returned wrong type. Have %T, want %s", test, testType.Name())
			}
			_ = testFnV.Call([]reflect.Value{reflect.ValueOf(t), reflect.ValueOf(test)})
		})
	}
}
