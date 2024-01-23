package main

import "github.com/amirhnajafiz/ghoster/pkg/logger"

func main() {
	l := logger.New().Trace("A")
	b := l.Trace("B")

	l.Error(nil, "A")
	b.Error(nil, "B")
}
