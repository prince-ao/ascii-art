package test

import (
	"asciify/internal/commands"
	"bytes"
	"io"
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	commands.HandleInfo("./assets/1.jpg")

	w.Close()

	var buf bytes.Buffer

	io.Copy(&buf, r)

	os.Stderr = old

	want := "Image Size: (960 x 1280)\n"
	got := buf.String()

	if got != want {
		t.Errorf("execpted: '%s'; got: '%s'", want, got)
	}
}
