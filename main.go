package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	engine "github.org/akrck02/nightfall/game"
	"github.org/akrck02/nightfall/sockets"
)

const (
	MAX_FPS = 3.0 // Target frames per second
	MAX_UPS = 60.0 // Target updates per second
)

var (
	running = true
	debug   = true
)

func main() {
	fmt.Println("Starting game engine...")

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
	fmt.Println("\nInterrupt received. Stopping game engine...")
	sockets.CloseConnections()
  running = false
	wg.Wait() // Wait for goroutines to finish

	fmt.Println("Game engine stopped.")
}

func updateLoop() {
	updateInterval := time.Second / time.Duration(MAX_UPS)
	lastUpdate := time.Now()

  ticker := time.NewTicker(time.Second) // Ticker to log UPS every second
	defer ticker.Stop()
 	updateCount := 0 // Counter for updates

	for running {
    select {
		  case <-ticker.C:
			  // Log updates per second
			  if debug {
			  	fmt.Printf("UPS: %d\n", updateCount)
			  }
			  updateCount = 0 // Reset the counter
		  default:
			  currentTime := time.Now()
			  elapsed := currentTime.Sub(lastUpdate)

			  if elapsed >= updateInterval {
				  engine.Update()
				  updateCount++       // Increment update counter
				  lastUpdate = currentTime
			  }

			  time.Sleep(time.Millisecond)
		}
	}
}

func drawLoop() {
	frameInterval := time.Second / time.Duration(MAX_FPS)
	lastFrame := time.Now()

  frameCount := 0 // Counter for frames
  fps := 0

	ticker := time.NewTicker(time.Second) // Ticker to log FPS every second
	defer ticker.Stop()

	for running {
    select {
		  case <-ticker.C:
			  // Log frames per second
			  fps = frameCount
			  frameCount = 0 // Reset the counter
		  default:
			  currentTime := time.Now()
			  elapsed := currentTime.Sub(lastFrame)

			  if elapsed >= frameInterval {
				  engine.Draw()
          println(fps)
				  frameCount++       // Increment frame counter
				  lastFrame = currentTime
			  }

			  time.Sleep(time.Millisecond)
		}
	}
}
