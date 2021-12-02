package main

import (
	"bufio"
	"io"
	"strconv"
)

func count(r io.Reader) int {
	first := true
	prev := 0
	c := 0
	for v := range read(r) {
		if first {
			// don't bother trying to count
			prev = v
			first = false
			continue
		}
		if v > prev {
			c++ // just so i could write c++
		}
		prev = v
	}
	return c
}

func read(r io.Reader) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		scn := bufio.NewScanner(r)
		for scn.Scan() {
			v, err := strconv.Atoi(scn.Text())
			if err != nil {
				return
			}
			ch <- v
		}
	}()

	return ch
}
