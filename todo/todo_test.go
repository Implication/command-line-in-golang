package todo_test

import (
	"os"
	"testing"

	"github.com/Implication/command-line-in-golang/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l[0].Task)
	}
}

func TestUpdate(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l[0].Task)
	}
	newTask := "Update Task"
	l.Update(1, newTask)
	if l[0].Task != newTask || l[0].Task == taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l[0].Task)
	}
}

func TestSearch(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	l.Add("Task 2")
	l.Add("Task 4")
	l.Add("Task 3")
	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l[0].Task)
	}

	if l.Search("Task 3") != true {
		t.Errorf("Search did not find task")
	}

	if l.Search("New Task") != true {
		t.Errorf("Search did not find task")
	}

	if l.Search("Update Task") == true {
		t.Errorf("Search unexpectedly found a task")
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l[0].Task)
	}

	l.Delete(1)
	if len(l) != 0 {
		t.Errorf("There should be no tasks in the list")
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)
	if l1[0].Task != taskName {
		t.Errorf("Expected %q got %q instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list to file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task. ", l1[0].Task, l2[0].Task)
	}
}
