package main

import "time"

type Line struct {
	Soccer   float64
	Football float64
	Baseball float64
}

func fetchLine(c chan Line) {
	for {
		time.Sleep(1 * time.Second)
		c <- Line{
			Soccer:   1.5,
			Football: 2.5,
			Baseball: 3.5,
		}
	}
}

func main() {
	c := make(chan Line)
	go fetchLine(c)

	for {
		select {
		case line := <-c:
			println(line.Soccer)
			println(line.Football)
			println(line.Baseball)
		}
	}
}
