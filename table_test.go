package testy

import "testing"

func TestTableTests(t *testing.T) {
	addFunc := func(a, b int) int {
		return a + b
	}
	type ttTest struct {
		a, b   int
		output int
	}
	table := &Table{
		"one": func(_ *testing.T) interface{} {
			return ttTest{a: 1, output: 1}
		},
		// "two": func(_ *testing.T) interface{} {
		// 	return 1
		// },
	}
	table.Run(t, func(t *testing.T, test ttTest) {
		output := addFunc(test.a, test.b)
		if output != test.output {
			t.Errorf("Expected %d, got %d\n", test.output, output)
		}
	})
}
