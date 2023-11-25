package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func sendSignal(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}

func KillProcess(pid int) error {
	return sendSignal(pid, syscall.SIGKILL)
}

func StopProcess(pid int) error {
	return sendSignal(pid, syscall.SIGSTOP)
}

func ContinueProcess(pid int) error {
	return sendSignal(pid, syscall.SIGCONT)
}

func GetProcessInfo(pid int) (*os.Process, error) {
	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}

	return process, nil

}

func GetProcessCmdline(pid int) ([]string, error) {
	cmdlineFile := fmt.Sprintf("/proc/%d/cmdline", pid)
	content, err := ioutil.ReadFile(cmdlineFile)
	if err != nil || len(content) == 0 {
		return nil, err
	}

	cmdline := strings.ReplaceAll(string(content), "\x00", " ")
	args := strings.Split(cmdline, " ")

	return args, nil
}

func GetParentPids(pid int) []int {
	parents := []int{}

	for {
		statFile := fmt.Sprintf("/proc/%d/stat", pid)
		content, err := ioutil.ReadFile(statFile)
		if err != nil || len(content) == 0 {
			break
		}

		stat := string(content)
		ppid := strings.Split(stat, " ")[3]
		parentPid, err := strconv.Atoi(ppid)

		if parentPid == 0 || err != nil {
			break
		}

		parents = append(parents, parentPid)
		pid = parentPid
	}

	return parents
}

func GetParentPid(pid int) int {
	parents := GetParentPids(pid)
	if len(parents) == 0 {
		return 0
	}

	return parents[0]
}
