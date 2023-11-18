package checker

import (
	"consecure/constant"
	"consecure/util/log"
	"consecure/util/process"
	"encoding/json"
	"os/exec"
)

type ImageChecker struct{}

type TrivyResult struct {
	Results []struct {
		Vulnerabilities []struct {
		} `json:"vulnerabilities"`
	} `json:"results"`
}

func NewImageChecker() *ImageChecker {
	return &ImageChecker{}
}

func (ic *ImageChecker) Check(event *constant.EngineEvent) bool {
	return ic.CheckImageVulnerabilities(event)
}

func (ic *ImageChecker) CheckImageVulnerabilities(event *constant.EngineEvent) bool {
	imageName := event.EngineMeta.Args[0]
	cmd := exec.Command("./trivy", "image", imageName, "--skip-db-update", "--format", "json")
	log.Debugln("Running", cmd.String())

	output, err := cmd.Output()
	if err != nil {
		ic.fatal(event, err)
	}

	var trivyResult TrivyResult
	err = json.Unmarshal(output, &trivyResult)
	if err != nil {
		ic.fatal(event, err)
	}

	hasVulnerabilities := false
	for _, result := range trivyResult.Results {
		if len(result.Vulnerabilities) > 0 {
			hasVulnerabilities = true
			break
		}
	}

	if hasVulnerabilities {
		log.Infof("Detected vulnerabilities in Image (%s).", imageName)
	}

	return hasVulnerabilities
}

func (ic *ImageChecker) fatal(event *constant.EngineEvent, err error) {
	if err != nil {
		process.KillProcess(event.Event.Pid)
		log.Fatal(err)
	}
}
