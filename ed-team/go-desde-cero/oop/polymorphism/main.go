package main

import "fmt"

type Compressor interface {
	Compress()
}

type Zip struct{}

type Gzip struct{}

type SevenZip struct{}

func (z Zip) Compress() {
	fmt.Println("Zip Compress")
}

func (gz Gzip) Compress() {
	fmt.Println("Gzip Compress")
}

func (z SevenZip) Compress() {
	fmt.Println("SevenZip Compress")
}

func Factory(method uint) Compressor {
	switch method {
	case 0:
		return Zip{}
	case 1:
		return Gzip{}
	case 2:
		return SevenZip{}
	default:
		return nil
	}
}

func main() {
	fmt.Println("Digit the compressor type:")
	fmt.Println("Options are: 1) Zip, 2) GZip, 3) 7Zip")
	var methodCompressor uint
	_, err := fmt.Scan(&methodCompressor)
	if err != nil {
		fmt.Println(err)
		return
	}
	compressMethod := Factory(methodCompressor)
	if compressMethod == nil {
		panic("CompressMethod is nil")
	}
	compressMethod.Compress()
}
