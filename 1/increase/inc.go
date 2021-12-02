package main

import (
	"bufio"
	"io"
	"strconv"
)

func agg(ch chan int, c int) chan int {
	ac := make(chan int)

	go func() {
		defer close(ac)
		buf := make([]int, c, c)
		c := 0
		sum := func(n []int) int {
			s := 0
			for i := range n {
				s += n[i]
			}
			return s
		}

		for v := range ch {
			buf[c%cap(buf)] = v
			c++
			if c < 3 {
				continue
			}
			ac <- sum(buf)
		}
	}()

	return ac
}

func count(ch chan int) int {
	first := true
	prev := 0
	c := 0
	for v := range ch {
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
