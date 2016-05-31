package fcgi

import "io"

func readToEOF(r io.Reader) {
	b := make([]byte, 1024)
	for {
		_, err := r.Read(b)
		if err == io.EOF {
			return
		}
	}
}
