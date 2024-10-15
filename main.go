package main

import (
	"fmt"
	"os"
	"strings"

	png "github.com/Maduki-tech/GoCode/decoder"
)

func main() {
	filename := os.Args
	if len(filename) < 2 {
		fmt.Println("Please provide a file name")
		return
	}

	if _, err := os.Stat(filename[1]); os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	readFile(filename[1])
}

func readFile(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fileNameAsArray := strings.Split(fileName, ".")
	fileType := fileNameAsArray[len(fileNameAsArray)-1]

	switch fileType {
	case "png":
		content, err := png.Decode(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(content)
		break
	default:
		fmt.Println("File not supported")
	}

}
