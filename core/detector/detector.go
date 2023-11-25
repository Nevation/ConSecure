package detector

import (
	"consecure/constant"
	"consecure/core/checker"
	"consecure/core/checker/runtime"
	"consecure/core/detector/engine"
	"consecure/util/log"
)

type Detector struct {
	engines        []engine.EngineStrategy
	runtimeChecker *runtime.RuntimeChecker
	imageChecker   *checker.Checker
}

func NewDetector() *Detector {
	return &Detector{
		engines:        engine.GetEngineStrategies(),
		runtimeChecker: runtime.NewRuntimeChecker(),
		imageChecker:   checker.NewChecker(),
	}
}

func (d *Detector) Detect(event *constant.Event) {
	go d.runRuntimeChecker(event)
	go d.runImageChecker(event)
}

func (d *Detector) createEngineEvent(event *constant.Event, meta *constant.EngineMeta) *constant.EngineEvent {
	return &constant.EngineEvent{
		Event:      event,
		EngineMeta: meta,
	}
}

func (d *Detector) runRuntimeChecker(event *constant.Event) {
	d.runtimeChecker.Check(event)
}

func (d *Detector) runImageChecker(event *constant.Event) {
	for _, engine := range d.engines {
		if engine.IsTargetEngine(event) {
			meta := engine.GetEngineMeta(event)

			if meta == nil {
				break
			}

			engineEvent := d.createEngineEvent(event, meta)
			log.Debugln("Detected Event", engineEvent.EngineMeta.Command, engineEvent.EngineMeta.Args)

			d.imageChecker.Check(engineEvent)
			break
		}
	}
}
