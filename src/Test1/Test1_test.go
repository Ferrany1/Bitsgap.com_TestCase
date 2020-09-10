package Test1

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	var test string
	for i := 0; i < 5; i++ {
		test += "123\n"
	}
	done := capture()
	Test1()
	capturedOutput, err := done()
	if err != nil {
	}
	if capturedOutput != test {
		t.Errorf("Wrong output: %s\nShould be: %s", capturedOutput, test)
	}
}

func capture() func() (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	done := make(chan error, 1)

	save := os.Stdout
	os.Stdout = w

	var buf strings.Builder

	go func() {
		_, err := io.Copy(&buf, r)
		err2 := r.Close()
		if err2 != nil {
			log.Fatalln(err2)
		}
		done <- err
	}()

	return func() (string, error) {
		os.Stdout = save
		err3 := w.Close()
		if err3 != nil {
			log.Fatalln(err3)
		}
		err := <-done
		return buf.String(), err
	}
}
