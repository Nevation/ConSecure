package detector

import (
	"consecure/constant"
	"consecure/core/detector/engine"
	"consecure/util/process"
)

type Detector struct {
	engines []engine.EngineStrategy
}

func NewDetector() *Detector {
	return &Detector{
		engines: engine.GetEngineStrategies(),
	}
}

func (d *Detector) Detect(event *constant.Event) {
	for _, engine := range d.engines {
		if engine.IsTargetEngine(event) {
			meta := engine.GetEngineMeta(event)

			if meta == nil {
				break
			}

			println("meta", meta)
			process.StopProcess(event.Pid)

			engineEvent := d.createEngineEvent(event, meta)

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
	// TODO
	println("run Checker", event, "image", event.EngineMeta.Args[0])
	// process.KillProcess(event.Event.Pid)
	process.ContinueProcess(event.Event.Pid)
}
