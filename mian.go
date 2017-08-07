package main

import (
	"os"
	"fmt"
	"strings"
	"bytes"
	"bufio"
)

func main() {
	prefix := os.Getenv("PREFIX")
	filename := os.Getenv("FILE_NAME")
	if filename == "" {
		filename = "default.ini"
	}
	if prefix == "" {
		prefix = ""
	}
	f, err := os.Create("./"+filename)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		key := pair[0];
		if len(key) > len(prefix) && strings.Compare(key[0:len(prefix)], prefix) == 0 {
			var buf bytes.Buffer
			buf.WriteString(strings.ToLower(key[len(prefix):]))
			buf.WriteString(" = ")
			buf.WriteString(pair[1])
			fmt.Fprintln(w, buf.String())

		}
	}
	w.Flush()
}
