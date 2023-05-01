package main

import "github.com/GAlexandrD/architecture-lab-3/scripts"

func main() {
	scripts.SendCommandToPainter(`white
	bgrect 0.25 0.25 0.75 0.75
	figure 0.5 0.5
	green
	figure 0.6 0.6
	update`)
}
