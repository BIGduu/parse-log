package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("logs/log/user_resource.log")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	size := info.Size()
	fmt.Println(size)
	lengthEveryPart := size >> 2
	println(lengthEveryPart)
	part, err := readPart(file, 4, lengthEveryPart)
	if err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println(string(part))

}

func readPart(file *os.File, part int64, length int64) ([]byte, error) {
	offset := part*length
	bytes := make([]byte,1024)
	_, err := file.ReadAt(bytes, int64(offset))
	return bytes, err
}
