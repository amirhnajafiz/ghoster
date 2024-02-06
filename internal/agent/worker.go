package agent

// Worker is a top process handler,
// and it gives two channels for communicating
// to a process
type Worker interface {
	GetStdin() chan interface{}
	GetStdout() chan interface{}
}
