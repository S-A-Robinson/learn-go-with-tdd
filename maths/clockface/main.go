package main

import (
	"os"
	"time"

	"learn-go-with-tdd/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
