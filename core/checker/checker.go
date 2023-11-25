package checker

import (
	"consecure/constant"
	"consecure/core/protector"
	"consecure/util/process"
)

type Checker struct {
	imageChecker *ImageChecker
	protector    *protector.Protector
}

func NewChecker() *Checker {
	return &Checker{
		imageChecker: NewImageChecker(),
	}
}

func (c *Checker) Check(event *constant.EngineEvent) error {
	process.StopProcess(event.Event.Pid)

	if event.EngineMeta.Target == "IMAGE" {
		c.CheckImage(event)
	}

	return nil
}

func (c *Checker) CheckImage(event *constant.EngineEvent) {
	if c.imageChecker.Check(event) {
		c.protector.Protect(event.Event)
	}
}
