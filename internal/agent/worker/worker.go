package worker

import "log"

type Worker struct{}

func (w Worker) Work(path string) error {
	log.Println(path)

	return nil
}
