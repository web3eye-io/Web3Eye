package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	<-sigchan
	os.Exit(1)
}
