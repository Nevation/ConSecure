package main

import (
	"consecure/core/watcher"
	"consecure/util"
	"log"
)

func main() {
	log.Println("Started Consecure project.")

	eventWatcher := watcher.NewEventWatcher(util.TraceDebugFilePath)
	go eventWatcher.Start()

	select {}
}
