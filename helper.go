package testy

import "testing"

func helper(t *testing.T) func() {
	return t.Helper
}
