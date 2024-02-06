package worker

import (
	"log"
)

type Worker struct {
	Channel chan string
}

func (w Worker) Work() {
	for {
		path := <-w.Channel

		log.Println(path)
	}
}
