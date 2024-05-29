package extractor

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

type Worker struct {
	wg sync.WaitGroup
}

func NewWorker() *Worker {
	return &Worker{}
}

func saveToTextFile(fileName string, prompt string) {
	promptData := []byte(prompt)
	sourceFileDir := filepath.Dir(fileName) + "//" + filepath.Base(fileName) + ".txt"
	fmt.Println(sourceFileDir)
	newFile, err := os.Create(sourceFileDir)
	if err != nil {
		return
	}

	newFile.Write(promptData)
}

func (worker *Worker) preprocessFile(fileName string) {
	defer worker.wg.Done()

	if filepath.Ext(fileName) != ".png" {
		return
	} else {
		data, err := ExtractDataFromImage(fileName)

		if err != nil {
			return
		}
		saveToTextFile(fileName, data)
	}
}

func DirTraverse(dirPath string) error {
	files := filepath.Dir(dirPath)

	worker := NewWorker()

	err := filepath.Walk(files, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		worker.wg.Add(1)
		go worker.preprocessFile(path)
		return nil
	})

	worker.wg.Wait()

	if err != nil {
		return err
	}

	return nil
}
