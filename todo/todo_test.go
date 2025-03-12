package todo_test

import (
	"fmt"
	"testing"

	"example.com/todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "NewTask"
	l.Add(taskName)

	// Check if the task name in the List corresponds with the assigned task name
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	// Add a task
	taskName := "My Task #1"
	l.Add(taskName)

	// Test if the task was added successfully
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}

	// Test if the task was completed
	if l[0].Done {
		t.Errorf("New Task should not be completed.")
	}

	// Mark as completed
	l.Complete(1)

	// Test if the task was successfully completed
	if !l[0].Done {
		t.Errorf("New Task should be completed.")
	}
}