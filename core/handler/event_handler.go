package handler

import (
	"consecure/constant"
	"consecure/core/detector"
)

type EventHandler struct {
	detector *detector.Detector
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
	go h.detector.Detect(event)
}
