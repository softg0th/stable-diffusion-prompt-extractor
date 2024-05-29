package main

import (
	"bufio"
	"extractor/extractor"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input path to the directory:")

	for {
		pathToDir, _ := reader.ReadString('\n')
		err := extractor.DirTraverse(pathToDir)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println("Success")
			break
		}
	}
}
