package notes_project

import (
	"encoding/json"
	"strings"
)

type saver interface {
	Save()
}

func buildFilenameFromTitle(title string) string {
	filename := strings.ReplaceAll(title, " ", "_")
	return strings.ToLower(filename) + ".json"
}

func formatContent(note *Note) []byte {
	data, err := json.Marshal(*note)

	if err != nil {
		return []byte("")
	}

	return data
}
