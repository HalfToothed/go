package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Note struct {
	Content string `json:"content"`
}

type NoteWrapper struct {
	Notes []Note `json:"noteWrapper"`
}

var Notes = []string{}

func addNote(note string) (result string) {

	file, err := os.ReadFile("notes.json")

	if err != nil {
		panic(err)
	}

	var savedNotes NoteWrapper
	if err := json.Unmarshal(file, &savedNotes); err != nil {
		fmt.Println("Error parsing json data", err)
	}

	newNote := Note{
		Content: note,
	}

	savedNotes.Notes = append(savedNotes.Notes, newNote)

	updatedData, err := json.MarshalIndent(savedNotes, "", " ")

	if err != nil {
		log.Fatalf("Error encoding JSON data: %v", err)
	}

	if err := os.WriteFile("notes.json", updatedData, 0644); err != nil {
		return "failed to add the note"
	} else {
		return note
	}
}

func listNote() {

	file, err := os.ReadFile("notes.json")

	if err != nil {
		panic(err)
	}

	var savedNotes NoteWrapper
	if err := json.Unmarshal(file, &savedNotes); err != nil {
		fmt.Println("Error parsing json data", err)
	}

	for i, notes := range savedNotes.Notes {
		fmt.Printf("%d: %s \n", i+1, notes.Content)
	}
}
