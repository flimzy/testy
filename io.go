package testy

import (
	"io"
	"os"
)

// Stdout redirects os.Stdout, and returns a reader to access the data that is
// sent to stdout, along with a cleanup function to restore normal functionality.
func Stdout() (io.Reader, func()) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	orig := os.Stdout
	os.Stdout = w
	return r, func() { os.Stdout = orig }
}

// Stderr redirects os.Stderr, and returns a reader to access the data that is
// sent to stdout, along with a cleanup function to restore normal functionality.
func Stderr() (io.Reader, func()) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	orig := os.Stderr
	os.Stderr = w
	return r, func() { os.Stderr = orig }
}

// Stdin redirects os.Stdin, and returns a writer to send data to stdin, along
// with a cleanup function to restore normal functionality.
func Stdin() (io.Writer, func()) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	orig := os.Stdin
	os.Stdin = r
	return w, func() { os.Stdin = orig }
}
