package runtime

import (
	"consecure/constant"
	"consecure/core/protector"
	"consecure/util/log"
	"consecure/util/process"
)

type RuntimeChecker struct {
	RunTimeContainerProcess *RuntimeContainerProcess
	protector               *protector.Protector
}

func NewRuntimeChecker() *RuntimeChecker {
	return &RuntimeChecker{
		RunTimeContainerProcess: NewRuntimeContainerProcess(),
	}
}

func (r *RuntimeChecker) isContainerProcess(pid int) bool {
	cmdLine, err := process.GetProcessCmdline(pid)

	if err != nil || len(cmdLine) < 1 {
		return false
	}

	if cmdLine[0] == "/usr/bin/containerd-shim-runc-v2" && cmdLine[1] == "-namespace" && cmdLine[3] == "moby" {
		go r.RunTimeContainerProcess.AddProcessInfo(pid, cmdLine)
		return false
	}

	parentPid := process.GetParentPid(pid)
	if r.RunTimeContainerProcess.IsContainerPid(pid, parentPid) {
		log.Debugln("Detected Runtime Event", pid, cmdLine)
		return true
	}

	return false
}

func (r *RuntimeChecker) Check(event *constant.Event) error {
	if !r.isContainerProcess(event.Pid) {
		return nil
	}

	log.Warningln("Detected runtime process that is not Entrypoint.", event.Pid)

	process.StopProcess(event.Pid)
	r.protector.Protect(event)

	return nil
}
