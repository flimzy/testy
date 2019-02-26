package testy

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestErrWriter(t *testing.T) {
	t.Run("0 bytes", func(t *testing.T) {
		w := ErrorWriter(0, errors.New("foo err"))
		_, err := fmt.Fprintln(w, "foo")
		Error(t, "foo err", err)
	})
	t.Run("5 bytes", func(t *testing.T) {
		w := ErrorWriter(5, errors.New("foo err"))
		c, err := fmt.Fprintln(w, "this is more than 5 bytes")
		if c != 5 {
			t.Errorf("Expected to write 5 bytes, wrote %d", c)
		}
		Error(t, "foo err", err)
	})
	t.Run("two writes", func(t *testing.T) {
		w := ErrorWriter(5, errors.New("foo err"))
		_, err := fmt.Fprintln(w, "xx")
		if err != nil {
			t.Errorf("Expected no error yet, got: %s", err)
		}
		c, err := fmt.Fprintln(w, "this is more than 5 bytes")
		if c != 2 {
			t.Errorf("Expected to write 5 bytes, wrote %d", c)
		}
		Error(t, "foo err", err)
	})
}

func TestErrReader(t *testing.T) {
	t.Run("0 bytes", func(t *testing.T) {
		r := ErrorReader("", errors.New("foo err"))
		_, err := ioutil.ReadAll(r)
		Error(t, "foo err", err)
	})
	t.Run("buffer", func(t *testing.T) {
		r := ErrorReader("foo text", errors.New("foo err"))
		c, err := ioutil.ReadAll(r)
		Error(t, "foo err", err)
		if string(c) != "foo text" {
			t.Errorf("Unexpected data returned: %s", string(c))
		}
	})
}
