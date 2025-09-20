package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath string, outputFilePath string) *FileManager {
	return &FileManager{inputFilePath, outputFilePath}
}

func (f *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(f.InputFilePath)
	if err != nil {
		return nil, errors.New("could not open file")
	}

	// this will be executed once the surrounding block or function finishes its execution
	defer func(file *os.File) { // 'defer' is used to defer tha operation
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (f *FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(f.OutputFilePath)
	if err != nil {
		return errors.New("could not create a file")
	}

	time.Sleep(3 * time.Second) // simulate delay

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("could not encode a file")
	}
	return nil
}
