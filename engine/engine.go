package engine

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/models"
)

var (
	screen image.RGBA              = *image.NewRGBA(image.Rect(0, 0, constants.ScreenWidth, constants.ScreenHeight))
	nodes  map[string]*models.Node = map[string]*models.Node{}
)

func loadResources() {
}

// Render the game frame and send to clients
func Draw() {
	screen = *image.NewRGBA(image.Rect(0, 0, constants.ScreenWidth, constants.ScreenHeight))
}

// Update the game
func Update(delta int) {
	for y := range constants.ScreenHeight {
		for x := range constants.ScreenWidth {
			screen.Set(x, y, color.RGBA{
				uint8(20),
				uint8(20),
				uint8(20),
				uint8(255),
			})
		}
	}

	// Draw()
}

// Add node to the game
func AddNode(node *models.Node) {
	nodes[node.Uuid] = node
}

// Get frame as byte array
func GetFrame() ([]byte, error) {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, &screen, nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
