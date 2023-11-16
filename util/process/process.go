package process

import (
	"fmt"
	"io/ioutil"
	"os"
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
	println("cmdline", cmdline, "cmdlineFile", cmdlineFile)
	args := strings.Split(cmdline, " ")

	return args, nil
}
