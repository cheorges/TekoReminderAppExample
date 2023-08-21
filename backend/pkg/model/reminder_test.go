package model

import (
	"testing"
)

func TestNewReminder(t *testing.T) {
	task := "Jumala"
	reminder := New(task)

	if reminder.Task != task {
		t.Errorf("Expected task to be %s, but got %s", task, reminder.Task)
	}

	if reminder.Done {
		t.Errorf("New reminder should not be marked as done")
	}
}

func TestToggleReminder(t *testing.T) {
	reminder := New("Jumala")

	reminder.Toggle()
	if !reminder.Done {
		t.Error("Toggle did not mark the reminder as done")
	}

	reminder.Toggle()
	if reminder.Done {
		t.Error("Toggle did not mark the reminder as not done")
	}
}
