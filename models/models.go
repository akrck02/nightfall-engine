package models

type Node struct {
	Uuid   string
	X      int
	Y      int
	Sprite rune
}

type (
	UpdateFunction func(delta int)
	DrawFunction   func()
)
