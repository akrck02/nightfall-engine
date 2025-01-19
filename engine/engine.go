package engine

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/models"
)

var (
	screen *image.RGBA = image.NewRGBA(
		image.Rect(0, 0, constants.ScreenWidth, constants.ScreenHeight),
	)
	nodes map[string]*models.Node = map[string]*models.Node{}
)

var spriteBuffer map[string]*image.Image = map[string]*image.Image{}

// Render the game frame and send to clients
func Draw() {
	for _, node := range nodes {
		if nil != node.Sprite {
			draw.Draw(screen, screen.Rect, *node.Sprite, image.Point{0, 0}, draw.Over)
		}
	}
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

	Draw()
}

// Add node to the game
func AddNode(node *models.Node) {
	nodes[node.Uuid] = node
}

// Get frame as byte array
func GetFrame() ([]byte, error) {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, screen.SubImage(screen.Rect), nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Load an sprite or getting from buffer
func LoadSprite(path string) *image.Image {
	hash := hashMd5(path)
	if nil != spriteBuffer[hash] {
		return spriteBuffer[hash]
	}

	file, err := os.Open(path)
	if nil != err {
		println(err.Error())
	}

	sprite, _, err := image.Decode(file)
	if err != nil {
		println(err.Error())
	}

	spriteBuffer[hash] = &sprite
	return &sprite
}

// Hash an string with md5
func hashMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
