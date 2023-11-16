package engine

import (
	"consecure/constant"
)

type EngineStrategy interface {
	IsTargetEngine(event *constant.Event) bool
	GetEngineMeta(event *constant.Event) *constant.EngineMeta
}
