package engine

import (
	"math/rand"
	"os"
	"os/exec"

	"github.org/akrck02/nightfall/constants"
	"github.org/akrck02/nightfall/sockets"
)

type RGBA64 struct {R, G, B, A uint16}
var screen [][]uint32 = [][]uint32{}

func Draw() {
  clear()
  sockets.SendFrameAsHtml(screen)
}

func Update() {
  
  screen = [][]uint32{}
  for i := 0; i < constants.Resolution[0] * constants.Resolution[1]; i++ {
    screen = append(screen, []uint32{
      uint32(rand.Intn(255)),
      uint32(rand.Intn(255)),
      uint32(rand.Intn(255)),
      uint32(rand.Intn(255)),
    })
  }
}

func clear() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}


