package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Trace("trace msg")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")
	log.Fatal("fatal msg")
	log.Panic("panic msg")
	//cmd.Execute()
}
