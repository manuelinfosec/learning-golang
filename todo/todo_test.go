package todo_test

import (
	"io/ioutil"
	"os"
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

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q", l[0].Task, tasks[0])
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected %q, got %q", 2, len(l))
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Erroor creating temp file: %s", err)

	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error savinig list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l2[0].Task, l1[0].Task)
	}
}
