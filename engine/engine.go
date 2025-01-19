package engine

import (
	"bytes"
	"fmt"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/models"
)

var (
	screen [][]rune
	nodes  map[string]*models.Node = map[string]*models.Node{}
)

// Render the game frame and send to clients
func Draw() {
	// Render the frame as ASCII
	var buffer bytes.Buffer
	for y := range screen {
		buffer.WriteString("\n")

		for _, rune := range screen[y] {
			buffer.WriteString(string(rune))
		}
	}

	fmt.Println(buffer.String())
}

// Update the game
func Update(delta int) {
	screen = make([][]rune, constants.ScreenHeight)
	for y := range screen {
		screen[y] = make([]rune, constants.ScreenWidth)
		for x := range screen[y] {
			screen[y][x] = ' '
		}
	}

	for _, node := range nodes {
		screen[node.Y][node.X] = node.Sprite
	}
}

// Add node to the game
func AddNode(node *models.Node) {
	nodes[node.Uuid] = node
}
