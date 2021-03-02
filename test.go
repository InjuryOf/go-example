package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("test.txt")
	if err == io.EOF {
		fmt.Println("read success")
	}
	fmt.Println(string(content))

	fileList, err := ioutil.ReadDir("io_example")
	if err == io.EOF {
		fmt.Println("read dir success")
	}
	for _, v := range fileList {
		fmt.Println("filename：", v.Size(), v.Name())
	}
	fmt.Println(os.TempDir())

	fmt.Println(strings.Title("hello world"))
}
