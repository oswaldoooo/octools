package toolsbox

import (
	"fmt"
	"strings"
)

const (
	template = "\r[%-101s] [%d%%]"
)

type Process struct {
	Pos        chan uint
	putchar    byte
	isfinished bool
}

func ProcessInit(putchar byte) *Process {
	return &Process{Pos: make(chan uint), putchar: putchar, isfinished: false}
}
func ProcessFinished(pro *Process) {
	if !pro.isfinished {
		pro.isfinished = true
	}
}

// 0 hidden 1 show
func ProcessRun(pro *Process, end *bool) {
	var process string
	for !pro.isfinished {
		select {
		case pos := <-pro.Pos:
			process = strings.Repeat(string(pro.putchar), int(pos)) + ">"
			fmt.Printf(template, process, pos)
		default:
			break
		}
	}
	fmt.Printf(template, strings.Repeat(string(pro.putchar), 100)+">", 100)
	// fmt.Print(" finished")
	*end = true
}
