package worker

import (
	"fmt"

	"github.com/amirhnajafiz/ghoster/pkg/enum"
	"github.com/amirhnajafiz/ghoster/pkg/utils"

	"github.com/google/uuid"
)

// Worker is a process manager for executing
// user projects by creating a new process
// and returning the process stdout.
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

// GetStdin of the worker.
func (w *Worker) GetStdin() chan interface{} {
	return w.stdin
}

// GetStdout of the worker.
func (w *Worker) GetStdout() chan interface{} {
	return w.stdout
}

// Work handles the logic of our
// processor.
func (w *Worker) Work() {
	for {
		data := <-w.stdin
		if data == enum.CodeDismiss {
			w.pipe <- 1

			return
		}

		// create source and dest addresses
		source := data.(string)
		dest := fmt.Sprintf("./data/%s/file.zip", uuid.NewString())

		// copy file to new path
		if err := utils.CopyFileContents(source, dest); err != nil {
			w.stdout <- enum.CodeFailure
			w.pipe <- 1

			return
		}

		// TODO: unzip the path
		// TODO: execute main.go
		// TODO: return output

		w.stdout <- data
	}
}
