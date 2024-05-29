package extractor

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func ExtractDataFromImage(imageName string) (string, error) {
	fmt.Println("start")
	imageFile, err := os.ReadFile(imageName)

	if err != nil {
		fmt.Printf("Somethig wrong with image %s!\n", imageName)
		return "", err
	}

	hexImageString := hex.EncodeToString(imageFile)
	startMarker := "74455874706172616d65746572"
	endMarker := "49444154"

	startIndex := strings.Index(hexImageString, startMarker)
	if startIndex == -1 {
		return "", nil
	}
	startIndex += len(startMarker)

	endIndex := strings.Index(hexImageString, endMarker)
	if endIndex == -1 {
		return "", nil
	}

	ordArray, err := hex.DecodeString(hexImageString[startIndex:endIndex])

	if err != nil {
		return "", nil
	}

	var runeArray []rune

	for _, ord := range ordArray {
		runeArray = append(runeArray, rune(ord))
	}

	prompt := string(runeArray)

	return prompt, nil
}
