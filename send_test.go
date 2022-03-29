package main

import (
	"bufio"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestReader struct{}
type TestWriter struct{}

func (tr *TestReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (tw *TestWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func TestSendFile(t *testing.T) {
	t.Parallel()

	var rw = bufio.NewReadWriter(bufio.NewReader(new(TestReader)), bufio.NewWriter(new(TestWriter)))
	assert.NoError(t, sendFile("testfile.txt", rw))
}
