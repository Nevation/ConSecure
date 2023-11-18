package protector

import (
	"consecure/constant"
	"consecure/util/log"
	"consecure/util/process"
)

type Protector struct{}

func NewProtector() *Protector {
	return &Protector{}
}

func (p *Protector) Protect(event *constant.EngineEvent) {
	process.KillProcess(event.Event.Pid)
	log.Infoln("Protector kill process", event.Event.Pid)
}
