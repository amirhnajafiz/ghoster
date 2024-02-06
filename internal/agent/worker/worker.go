package worker

import (
	"log"
)

type Worker struct {
	Terminate chan bool
	Pipe      chan int
	Channel   chan string
}

func (w Worker) Work() {
	for {
		select {
		case path := <-w.Channel:
			log.Println(path)
		case <-w.Terminate:
			w.Pipe <- 1

			return
		}
	}
}
