package agent

type Worker interface {
	GetStdin() chan interface{}
	GetStdout() chan interface{}
}
