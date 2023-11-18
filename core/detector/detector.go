package detector

import (
	"consecure/constant"
	"consecure/core/checker"
	"consecure/core/detector/engine"
	"consecure/util/log"
	"consecure/util/process"
)

type Detector struct {
	engines []engine.EngineStrategy
	checker *checker.Checker
}

func NewDetector() *Detector {
	return &Detector{
		engines: engine.GetEngineStrategies(),
		checker: checker.NewChecker(),
	}
}

func (d *Detector) Detect(event *constant.Event) {
	for _, engine := range d.engines {
		if engine.IsTargetEngine(event) {
			meta := engine.GetEngineMeta(event)

			if meta == nil {
				break
			}

			engineEvent := d.createEngineEvent(event, meta)
			log.Debugln("Detected Event", engineEvent.EngineMeta.Command, engineEvent.EngineMeta.Args)

			go d.runChecker(engineEvent)
			break
		}
	}
}

func (d *Detector) createEngineEvent(event *constant.Event, meta *constant.EngineMeta) *constant.EngineEvent {
	return &constant.EngineEvent{
		Event:      event,
		EngineMeta: meta,
	}
}

func (d *Detector) runChecker(event *constant.EngineEvent) {
	process.StopProcess(event.Event.Pid)

	log.Debugln(
		"run Checker",
		"pid", event.Event.Pid,
		"target", event.EngineMeta.Target,
		"command", event.EngineMeta.Command,
		"args", event.EngineMeta.Args,
	)

	d.checker.Check(event)
}
