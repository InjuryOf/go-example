package io_example

import (
	"bufio"
	"io"
)

// reader io
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

// write io
func WriteFrom(file io.Writer, data string) (int, error) {
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	return writer.WriteString(data)
}
