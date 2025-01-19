package game

import (
	"github.org/akrck02/nightfall/engine"
	"github.org/akrck02/nightfall/models"
)

func LoadMap() {
	println("Loading map!")

	player := &models.Node{
		Uuid:   "Shadow",
		X:      0,
		Y:      0,
		Sprite: engine.LoadSprite("test.png"),
	}
	engine.AddNode(player)
}
