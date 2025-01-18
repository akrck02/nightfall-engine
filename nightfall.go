package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/engine"
	"github.org/akrck02/nightfall/sockets"
	"github.org/akrck02/nightfall/stats"
)

var running = true

func StartMainLoop() {
	// Channel to listen for interrupt signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	// Wait group to ensure goroutines are properly managed
	var wg sync.WaitGroup
	wg.Add(2)

	// Start the update loop
	go func() {
		defer wg.Done()
		updateLoop()
	}()

	// Start the draw loop
	go func() {
		defer wg.Done()
		drawLoop()
	}()

	go sockets.Start()

	// Wait for the signal to stop the game
	<-signalChannel
	running = false
	wg.Wait() // Wait for goroutines to finish
}

// This is the logic for the update loop
func updateLoop() {
	updateInterval := time.Second / time.Duration(constants.MAX_UPS)
	lastUpdate := time.Now()

	ticker := time.NewTicker(time.Second) // Ticker to log UPS every second
	defer ticker.Stop()
	updateCount := 0 // Counter for updates
	stats.Ups = 0

	for running {
		select {
		case <-ticker.C:
			stats.Ups = updateCount
			updateCount = 0 // Reset the counter
		default:
			currentTime := time.Now()
			elapsed := currentTime.Sub(lastUpdate)

			if elapsed >= updateInterval {
				engine.Update(int(elapsed.Nanoseconds()))
				updateCount++ // Increment update counter
				lastUpdate = currentTime
			}

			time.Sleep(time.Millisecond)
		}
	}
}

// This is the logic for the draw loop
func drawLoop() {
	frameInterval := time.Second / time.Duration(constants.MAX_FPS)
	lastFrame := time.Now()

	frameCount := 0 // Counter for frames
	stats.Fps = 0

	ticker := time.NewTicker(time.Second) // Ticker to log FPS every second
	defer ticker.Stop()

	for running {
		select {
		case <-ticker.C:
			stats.Fps = frameCount
			frameCount = 0 // Reset the counter
		default:
			currentTime := time.Now()
			elapsed := currentTime.Sub(lastFrame)

			if elapsed >= frameInterval {
				engine.Draw()
				frameCount++ // Increment frame counter
				lastFrame = currentTime
			}

			time.Sleep(time.Millisecond)
		}
	}
}
