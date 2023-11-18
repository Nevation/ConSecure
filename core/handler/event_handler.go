package handler

import (
	"consecure/constant"
	"consecure/core/detector"
)

type EventHandler struct {
	detector *detector.Detector
	pids     []int
}

var eventHandlerInstance *EventHandler

func GetEventHandler() *EventHandler {
	if eventHandlerInstance == nil {
		eventHandlerInstance = &EventHandler{
			detector: detector.NewDetector(),
		}
	}

	return eventHandlerInstance
}

func (h *EventHandler) HandleEvent(event *constant.Event) {
	if contains(h.pids, event.Pid) {
		return
	}

	h.pids = append(h.pids, event.Pid)
	go h.detector.Detect(event)
}

func contains(s []int, elem int) bool {
	for _, a := range s {
		if a == elem {
			return true
		}
	}
	return false
}
