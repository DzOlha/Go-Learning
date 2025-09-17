package notes_project

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetNoteData() *Note {
	title := getInputData("Title: ")
	content := getInputData("Content: ")

	return NewNote(title, content)
}

func getInputData(prompt string) string {
	var data string

	for data == "" {
		data = showInputField(prompt)
	}

	return data
}

func showInputField(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
