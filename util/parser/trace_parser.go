package parser

import (
	"consecure/constant"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseTraceLine(line string) (*constant.Event, error) {
	r := regexp.MustCompile(`(?P<comm>\w+)-(?P<pid>\d+)\s+\[\d+\]\s+\.{5}\s+(?P<timestamp>\d+\.\d+):\s+(?P<event>\w+):`)
	match := r.FindStringSubmatch(line)
	if len(match) == 0 {
		return nil, nil
	}

	pid, err := strconv.Atoi(match[2])
	if err != nil {
		return nil, err
	}

	command := strings.Split(match[1], "-")[0]
	event := &constant.Event{
		Pid:       pid,
		Timestamp: time.Now(),
		Command:   command,
		EventName: match[4],
	}

	return event, nil
}
