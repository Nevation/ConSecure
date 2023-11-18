package engine

import (
	"consecure/constant"
	"consecure/util/process"
)

type DockerEngine struct {
}

func NewDockerEngine() *DockerEngine {
	return &DockerEngine{}
}

func (e *DockerEngine) IsTargetEngine(event *constant.Event) bool {
	return event.Command == "docker"
}

func (e *DockerEngine) GetEngineMeta(event *constant.Event) *constant.EngineMeta {
	cmdLines, err := process.GetProcessCmdline(event.Pid)

	if err != nil {
		return nil
	}

	target := ""

	if e.isTargetTypeImage(cmdLines) {
		target = "IMAGE"
	} else {
		return nil
	}

	return &constant.EngineMeta{
		Target:  target,
		Command: cmdLines[1],
		Args:    cmdLines[2:],
	}
}

func (e *DockerEngine) isTargetTypeImage(cmdLines []string) bool {
	mainCmd := cmdLines[1]

	if mainCmd == "pull" {
		return true
	}

	return mainCmd == "run"
}
