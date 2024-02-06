package worker

import "log"

type Worker struct {
	Pipe   chan string
	Stdout chan string
}

func (w Worker) Work() {
	for {
		path := <-w.Pipe

		log.Println(path)
	}
}
