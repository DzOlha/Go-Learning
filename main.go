package main

import notes_project "go-learning/main/05-structs_custom-types/notes-project"

func main() {
	//structs.RunTheApp()
	//structs.AliasExample()
	note := notes_project.GetNoteData()
	note.Print()
	note.Save()
}
