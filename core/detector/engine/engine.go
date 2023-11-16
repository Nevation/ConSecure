package engine

import (
	"consecure/constant"
)

type Engine struct {
	strategy EngineStrategy
}

func NewEngine(strategy EngineStrategy) *Engine {
	return &Engine{
		strategy: strategy,
	}
}

func GetEngineStrategies() []EngineStrategy {
	return []EngineStrategy{
		NewDockerEngine(),
	}
}

func (e *Engine) SetStrategy(strategy EngineStrategy) {
	e.strategy = strategy
}

func (e *Engine) IsTargetEngine(event *constant.Event) bool {
	return e.strategy.IsTargetEngine(event)
}

func (e *Engine) GetEngineMeta(event *constant.Event) *constant.EngineMeta {
	return e.strategy.GetEngineMeta(event)
}
