package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// /Users/yorick/Downloads/auth-resource.jar
	args := os.Args
	if args == nil || len(args) != 3 {
		fmt.Println("useage : xxx srcFile destFile")
		return
	}
	srcPath := args[1]
	destPath := args[2]
	fmt.Printf("srcPath=%s,destPath=%s\n", srcPath, destPath)
	if srcPath == destPath {
		fmt.Println("the same file")
		return
	}
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("srcFile open err: ", err)
		return
	}
	defer srcFile.Close()
	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Println("destFile create err: ", err)
		return
	}
	defer destFile.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		if n == 0 {
			fmt.Println("read finish!")
			break
		}
		destFile.Write(buf[:n])
	}
}
