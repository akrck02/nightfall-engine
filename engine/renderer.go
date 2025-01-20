package engine

func ScreenY(y int, height int, maxHeight int) int {
	newY := height - y

	if newY <= 0 {
		return 0
	}

	if newY+height >= maxHeight {
		return maxHeight - height
	}

	return newY
}

func ScreenX(x int, width int, maxWidth int) int {
	newX := x

	if newX <= 0 {
		return 0
	}

	if newX+width >= maxWidth {
		return maxWidth - width
	}

	return newX

}
