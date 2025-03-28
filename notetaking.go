package main

import (
	"fmt"
	"net/http"
	"html/template"
)

// Note structure
type Note struct {
	Content string
}

// Slice to hold notes
var notes []Note

// Handler to display notes
func notesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Shared Notes</title>
</head>
<body>
    <h1>Shared Notes</h1>
    <form method="POST" action="/add-note">
        <textarea name="note" placeholder="Write a note..." rows="4" cols="50"></textarea><br><br>
        <button type="submit">Add Note</button>
    </form>
    <div>
        <h2>Notes:</h2>
        {{range .}}
            <div style="margin-bottom: 10px; padding: 10px; background-color: #f9f9f9;">
                {{.Content}}
            </div>
        {{end}}
    </div>
</body>
</html>
`)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, notes)
}

// Handler to add a note
func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		note := r.FormValue("note")
		if note != "" {
			notes = append(notes, Note{Content: note})
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/", notesHandler)
	http.HandleFunc("/add-note", addNoteHandler)

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
