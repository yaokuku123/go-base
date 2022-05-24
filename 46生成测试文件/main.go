package main

import (
	"bytes"
	"os"
)

const fileName = "/Users/yorick/Downloads/"
const con_KB = 1024 * 1024

func main() {

	file, _ := os.Create(fileName + "test_4.5M.txt")
	content := bytes.Repeat([]byte("1"), 4.5*con_KB)
	file.Write(content)
	file.Close()

}
