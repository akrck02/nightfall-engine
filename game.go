package main

import (
	"github.com/google/uuid"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/engine"
	"github.org/akrck02/nightfall/models"
)

func main() {
	const playerSprite = 'ƒ'
	const enemySprite = 'Ø'

	player := models.Node{
		Uuid:   uuid.New().String(),
		X:      5,
		Y:      3,
		Sprite: playerSprite,
	}

	rightMovement := true

	engine.AddNode(&player)
	engine.AddUpdateFunction(func(delta int) {
		if rightMovement {
			player.X += 1
		} else {
			player.X -= 1
		}

		if player.X >= constants.ScreenWidth-2 {
			player.X = constants.ScreenWidth - 2
			rightMovement = false
		} else if player.X < 1 {
			player.X = 1
			rightMovement = true
		}
	})

	setLevelWalls()
	StartMainLoop()
}

func setLevelWalls() {
	const wall = '█'
	for i := 0; i < constants.ScreenWidth; i++ {
		engine.AddNode(&models.Node{
			Uuid:   uuid.New().String(),
			X:      i,
			Y:      0,
			Sprite: wall,
		})

		engine.AddNode(&models.Node{
			Uuid:   uuid.New().String(),
			X:      i,
			Y:      constants.ScreenHeight - 1,
			Sprite: wall,
		})
	}

	for i := 0; i < constants.ScreenHeight; i++ {

		engine.AddNode(&models.Node{
			Uuid:   uuid.New().String(),
			X:      0,
			Y:      i,
			Sprite: wall,
		})

		engine.AddNode(&models.Node{
			Uuid:   uuid.New().String(),
			X:      constants.ScreenWidth - 1,
			Y:      i,
			Sprite: wall,
		})
	}
}
