package main

import (
	"github.com/pollen5/minori"
)

var logger = minori.GetLogger("app")

func main() {
	logger.Info("Booting up...")
	logger.Warn("Hey, That's deprecated!")
	logger.Debug("Debugging the server.")
	logger.Error("Something failed.")

	// Recover the panic so we can also do a final view of the fatal function.
	defer func() {
		if err := recover(); err != nil {
			logger.Fatal("Something blew up.")
		}
	}()

	logger.Panic("Something panicked.")
}
