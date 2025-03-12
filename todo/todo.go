package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item structs represent a single task in the todo list
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add adds a task to the todo list
func (l *List) Add(task string) {
	// Append a new item to the list pointed to
	*l = append(*l, item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	})
}

// Complete method marks a ToDo item as completed by
// setting Done = true and ComopleedAt to the current time
func (l *List) Complete(idx int) error {
	ls := *l // Copy the slice to a local variable
	// Implement check to see if the index is valid
	if idx <= 0 || idx > len(ls) {
		return fmt.Errorf("item %d does not exist", idx)
	}

	// Adjusting index for 0 based index
	ls[idx-1].Done = true
	ls[idx-1].CompletedAt = time.Now()

	return nil // Return nil to indicate success
}

// Delete method removes a ToDo item from the list
func (l *List) Delete(idx int) error {
	ls := *l // Copy the slice to a local variable
	// Implement check to see if the index is valid
	if idx < 0 || idx >= len(ls) {
		return fmt.Errorf("item %d does not exist", idx)
	}

	// Getting the items from the beginning of the list and up to the index
	// and appending the items from the index to the end of the list
	*l = append(ls[:idx-1], ls[idx:]...)

	return nil
}

// String method retuns a string representation of the list
func (l *List) String() (string, error) {
	ls := *l

	// Implement check to see if any tasks exists
	if len(ls) == 0 {
		return "", fmt.Errorf("no tasks exists in the list")
	}

	// Initialize a string to hold the list of tasks
	items := ""

	// Loop through the list and add each task to the `items` string
	for item := range ls {
		items += fmt.Sprintf("%d. %s\n", item+1, ls[item].Task)
	}

	return items, nil
}

// Save method encodes the List as JSSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	// Serialaize the list to JSON
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	// Makes multiple system calls
	return os.WriteFile(filename, js, 0644)
}

// Get method reads a JSON file and decodes it into a List
func (l *List) Get(filename string) error {
	// Read the file
	file, err := os.ReadFile(filename)
	if err != nil {
		// Check if the file does not exist
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
	}

	// Check if the file is empty
	if len(file) == 0 {
		return nil
	}

	// Deserialize the JSON file into the List
	return json.Unmarshal(file, l)
}
