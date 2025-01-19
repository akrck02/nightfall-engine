package models

import "image"

type Node struct {
	Uuid   string
	X      int
	Y      int
	Sprite *image.Image
}
