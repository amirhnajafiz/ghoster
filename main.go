package main

import "github.com/amirhnajafiz/ghoster/pkg/logger"

func main() {
	l := logger.New(logger.ErrorLevel).Trace("A")
	b := l.Trace("B")

	l.Error(nil, "m", "A")
	b.Error(nil, "m", "B")

	c := b.Trace("C")

	c.Info(nil, "m", "C")
	c.Error(nil, "m", "C", "d", "B")
}
