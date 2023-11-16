package constant

import "time"

type Event struct {
	Pid       int
	Timestamp time.Time
	Command   string
	EventName string
}

type EngineMeta struct {
	Target  string
	Command string
	Args    []string
}

type EngineEvent struct {
	EngineMeta *EngineMeta
	Event      *Event
}
