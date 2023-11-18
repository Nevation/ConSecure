package main

import (
	"consecure/core/watcher"
	"consecure/util"
	"consecure/util/log"
)

func main() {
	log.Infoln("Started Consecure project.")

	eventWatcher := watcher.NewEventWatcher(util.TraceDebugFilePath)
	go eventWatcher.Start()

	select {}
}
