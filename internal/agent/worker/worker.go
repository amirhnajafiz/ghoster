package worker

type Worker struct {
	stdin  chan interface{}
	stdout chan interface{}
	pipe   chan int
}

func New(pipe chan int) *Worker {
	return &Worker{
		stdin:  make(chan interface{}),
		stdout: make(chan interface{}),
		pipe:   pipe,
	}
}

func (w *Worker) GetStdin() chan interface{} {
	return w.stdin
}

func (w *Worker) GetStdout() chan interface{} {
	return w.stdout
}

func (w *Worker) Work() {
	for {
		data := <-w.stdin
		if data == "DONE" {
			w.pipe <- 1

			return
		}

		w.stdout <- data
	}
}
