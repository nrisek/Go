package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := getArguments() // go run reader.go file.txt
	file := openFile(filePath)
	bs := make([]byte, 32*1024)
	file.Read(bs)
	fmt.Println(string(bs))
}

func getArguments() string {
	arguments := os.Args
	return arguments[1]
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return file
}
