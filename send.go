package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// declare chunk size
const bufferSize = 1e3

func sendFile(name string, rw *bufio.ReadWriter) error {

	// open file
	var file, err = os.Open(strings.TrimSpace(name))
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("File does not exists: %w", err)
	} else if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	// write the header: just filename for now
	_, err = rw.WriteString(strings.TrimSpace(name) + "\n")
	if err != nil {
		return fmt.Errorf("Error writing header to stream: %w", err)
	}

	// get some info for the file
	var stat, _ = file.Stat()

	// create buffer
	var buffer = make([]byte, bufferSize)

	var totalBytesReadFromFile int
	var totalBytesSent int
	for {
		// read content to buffer
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		totalBytesReadFromFile += bytesRead

		// handle the last cycle through the loop, the number of bytes read
		// from the file is likely to be less than the buffer size so we need
		// to remove the end of the buffer
		if bytesRead < bufferSize {
			buffer = buffer[:bytesRead]
		}

		// SEND IT
		_, err = rw.Write(buffer)
		if err != nil {
			return fmt.Errorf("Error writing payload to stream: %w", err)
		}
		totalBytesSent += len(buffer)

		// reset
		buffer = make([]byte, bufferSize)
	}

	// sanity check that we read everything correctly
	if int64(totalBytesReadFromFile) != stat.Size() {
		return fmt.Errorf("bytes read from file [%d] does not equal the file size [%d]", totalBytesReadFromFile, stat.Size())
	} else if int64(totalBytesReadFromFile) != int64(totalBytesSent) {
		return fmt.Errorf("bytes sent [%d] does not equal the file size [%d]", totalBytesSent, stat.Size())
	}

	return nil
}
