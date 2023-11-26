package utils

import (
	"bytes"
	"io"
)

// Конвертация []byte в io.Readcloser
func ConvertByteToReader(b []byte) io.ReadCloser {
	reader := bytes.NewReader(b)
	return io.NopCloser(reader)
}

// Конвертация io.Readcloser в []byte
func ConvertReaderToByte(r io.ReadCloser) ([]byte, error) {
	defer r.Close()
	return io.ReadAll(r)
}
