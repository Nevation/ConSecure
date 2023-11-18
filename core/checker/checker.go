package checker

import (
	"consecure/constant"
	"consecure/core/protector"
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
	if event.EngineMeta.Target == "IMAGE" {
		c.CheckImage(event)
	}

	return nil
}

func (c *Checker) CheckImage(event *constant.EngineEvent) {
	if c.imageChecker.Check(event) {
		c.protector.Protect(event)
	}
}
