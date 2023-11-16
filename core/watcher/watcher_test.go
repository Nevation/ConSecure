package watcher

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestEventWatcher_Start(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write some lines to the file
	lines := []string{"line 1", "line 2", "line 3"}
	for _, line := range lines {
		if _, err := tmpfile.WriteString(line + "\n"); err != nil {
			t.Fatal(err)
		}
	}

	// Create a new event watcher for the temporary file
	watcher := NewEventWatcher(tmpfile.Name())

	// Start the event watcher in a separate goroutine
	go watcher.Start()

	// Wait for the watcher to start
	time.Sleep(100 * time.Millisecond)

	// Append some new lines to the file
	newLines := []string{"line 4", "line 5"}
	for _, line := range newLines {
		if _, err := tmpfile.WriteString(line + "\n"); err != nil {
			t.Fatal(err)
		}
	}

	// Wait for the watcher to detect the new lines
	time.Sleep(1000 * time.Millisecond)

	// Check that the watcher detected the new lines
	expectedEvents := len(newLines)
	actualEvents := len(watcher.getFileLines()) - len(lines)
	if actualEvents != expectedEvents {
		t.Errorf("Expected %d events, but got %d", expectedEvents, actualEvents)
	}
}
