package main

import (
	"fmt"
	"math"
	"time"

	"github.com/GAlexandrD/architecture-lab-3/scripts"
)

func MoveOnCircle(a int, r float64) {
	x := r * math.Cos(float64(a))
	y := r * math.Sin(float64(a))
	scripts.SendCommandToPainter(fmt.Sprintf("move %f %f", x, y))
	scripts.SendCommandToPainter("update")
}

func main() {
	scripts.SendCommandToPainter("green")
	scripts.SendCommandToPainter("figure 0.4 0.2")
	var a = 0
	var r float64 = 0.2
	for {
		MoveOnCircle(a, r)
		time.Sleep(200 * time.Millisecond)
		a++
	}
}
