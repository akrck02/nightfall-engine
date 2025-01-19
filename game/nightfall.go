package game

import (
	"time"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/engine"
	"github.org/akrck02/nightfall/stats"
)

var running = true

// Start logic
func Start() {
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
			// println(fmt.Sprintf("UPS: %d", stats.Ups))
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

// Stop logic
func Stop() {
	running = false
}
