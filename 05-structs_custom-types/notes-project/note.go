package notes_project

import (
	"fmt"
	"os"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNote(title, content string) *Note {
	return &Note{
		Title:   title,
		Content: content,
	}
}

func (note *Note) Print() {
	fmt.Println(note.Title + ": " + note.Content)
}

func (note *Note) Save() {
	filename := buildFilenameFromTitle(note.Title)
	err := os.WriteFile(filename, formatContent(note), 0644)
	if err != nil {
		return
	}
}
