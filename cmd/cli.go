package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
    // Load in the document from standard in
    doc, err := loadInDoc()
	if err != nil {
		panic(err)
	}
    slog.Info("choose a path")
    // Provide option to choose a component
    paths := doc.Paths.InMatchingOrder()
    for idx, path := range paths {
        fmt.Printf("%d. %s\n", idx, path)
    }
}

func loadInDoc() (*openapi3.T, error) {
	// Read in the open api spec.
	buf := bytes.NewBuffer([]byte{})
	reader := bufio.NewReader(os.Stdin)
	for {
		lineContent, err := reader.ReadString('\n')
		if err == io.EOF {
			slog.Info("found end of file")
			break
		}
		// Write contents to a buffer.
		slog.Info("Reading line of file", "line", lineContent)
		buf.WriteString(lineContent)
	}
	loader := openapi3.NewLoader()
    return loader.LoadFromData(buf.Bytes())
}
