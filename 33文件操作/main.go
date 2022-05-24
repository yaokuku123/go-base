package main

import (
	"fmt"
	"os"
)

func writeFile() {
	f, err := os.Create("./test-file.txt")
	if err != nil {
		fmt.Println("file err: ", err)
		return
	}
	defer f.Close()
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("%s:%d\n", "hello go", i)
		f.WriteString(str)
		f.Write([]byte("#####\n"))
	}
}

func ReadFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file open err: ", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			fmt.Println("read finish,err: ", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	//writeFile()
	ReadFile("./test-file.txt")
}
