package watcher

import (
	"bufio"
	"consecure/constant"
	"consecure/core/handler"
	"consecure/util/parser"
	"os"
	"time"
)

type EventWatcher struct {
	filePath     string
	watcherPid   int
	eventHandler *handler.EventHandler
}

func NewEventWatcher(filePath string) *EventWatcher {
	return &EventWatcher{
		filePath:     filePath,
		watcherPid:   os.Getpid(),
		eventHandler: handler.GetEventHandler(),
	}
}

func (w *EventWatcher) getFileLines() []string {
	file, err := os.Open(w.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (w *EventWatcher) Start() {
	lastLineNumber := len(w.getFileLines())

	for {
		time.Sleep(100 * time.Millisecond)

		nowFileData := w.getFileLines()
		nowLineNumber := len(nowFileData)

		if lastLineNumber < nowLineNumber {
			newLines := nowFileData[lastLineNumber:]
			for _, newLine := range newLines {
				event, err := parser.ParseTraceLine(newLine)

				if err != nil || !w.filterEvent(event) {
					continue
				}

				if event.Pid == w.watcherPid {
					continue
				}

				w.eventHandler.HandleEvent(event)
			}
		}

		lastLineNumber = nowLineNumber
	}
}

func (w *EventWatcher) filterEvent(event *constant.Event) bool {
	if event == nil {
		return false
	}

	if event.Pid == w.watcherPid {
		return false
	}

	if event.EventName != "sched_process_fork" {
		return false
	}

	return true
}
