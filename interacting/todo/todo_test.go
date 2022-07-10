package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"ooi.cc/interacting/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, and got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, and got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Error("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Error("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q but got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected %d length but got %d instead", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q but got %q instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q but got %q instead", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating tempfile %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving to file %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match Task %q", l2[0].Task, l1[0].Task)
	}
}
