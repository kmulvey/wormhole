package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func readData(rw *bufio.ReadWriter) {
	for {
		var fileName, err = rw.ReadString('\n')
		if err != nil {
			log.Error("Error reading from buffer: ", err)
			return
		}
		if fileName == "" {
			return
		}

		fileName = strings.TrimSpace(strings.TrimSuffix(fileName, "\n"))
		f, err := os.Create(filepath.Base(fileName))
		if err != nil {
			log.Error("Error creating file: ", err)
			return
		}
		defer f.Close()

		_, err = rw.WriteTo(f)
		if err != nil {
			log.Error("Error writing file: ", err)
			return
		}
	}
}
