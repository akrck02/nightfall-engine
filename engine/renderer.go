package engine

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearCanvas() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	if nil == cmd {
		return
	}

	if cmd.Err != nil {
		println(cmd.Err.Error())
		return
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}
