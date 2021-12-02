package main

import (
	"bufio"
	"fmt"
	"io"
)

// returns an x and y position -- where x represents our horizontal movement
// after commands and y is our vertical movement after commands -- after
// parsing input.
func parse(r io.Reader) (int, int) {
	scn := bufio.NewScanner(r)
	var x, y int
	for scn.Scan() {

		var cmd string
		var v int
		fmt.Sscanf(scn.Text(), "%s %d", &cmd, &v)
		switch cmd {
		case "forward":
			x += v
		case "down":
			y += v
		case "up":
			y -= v
		}
	}
	return x, y
}
