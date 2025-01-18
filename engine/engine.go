package engine

import (
	"bytes"
	"fmt"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/models"
	"github.org/akrck02/nightfall/stats"
)

var (
	screen          [][]rune
	nodes           map[string]*models.Node = map[string]*models.Node{}
	updateFunctions []models.UpdateFunction = []models.UpdateFunction{}
)

// Render the game frame and send to clients
func Draw() {
	ClearCanvas()

	if constants.DEBUG_MODE {
		fmt.Println(fmt.Sprintf("UPS: %d | FPS: %d", stats.Ups+1, stats.Fps+1))
	}

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
	for _, fn := range updateFunctions {
		fn(delta)
	}

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

func AddUpdateFunction(updateFunction models.UpdateFunction) {
	updateFunctions = append(updateFunctions, updateFunction)
}
