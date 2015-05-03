package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (a MyReader) Read (b []byte) (c int, err error){
	b[0]= 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}