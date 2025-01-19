package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.org/akrck02/nightfall/discord"
	"github.org/akrck02/nightfall/game"
)

func main() {
	// Wait group to ensure goroutines are properly managed
	var wg sync.WaitGroup
	wg.Add(2)

	// Start the game
	go func() {
		defer wg.Done()
		game.Start()
	}()

	// Start discord server
	go func() {
		defer wg.Done()
		discord.Start()
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Nightfall is now running.  Press CTRL-C to exit.")
	sessionChannel := make(chan os.Signal, 1)
	signal.Notify(sessionChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sessionChannel

	// Cleanly close down the Discord session.
	discord.Stop()
	game.Stop()
}
