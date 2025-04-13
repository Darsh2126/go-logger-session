package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-session-logger/logger"
)

func main() {
	fmt.Println("🟢 Session started. Press Ctrl+C to stop...")

	start := time.Now()

	// Wait for Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	end := time.Now()
	duration := end.Sub(start)

	fmt.Println("🛑 Session stopped.")
	logger.WriteSession(start, end, duration)
}
