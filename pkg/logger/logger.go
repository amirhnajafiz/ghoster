package logger

type Logger interface {
	Error(err error, params ...interface{})
	Info(err error, params ...interface{})
	Debug(err error, params ...interface{})
	Trace(label string) Logger
}

func New() Logger {
	return &logger{}
}

type logger struct {
	tracer string
}

func (l logger) Error(err error, params ...interface{}) {

}

func (l logger) Info(err error, params ...interface{}) {

}

func (l logger) Debug(err error, params ...interface{}) {

}

func (l logger) Trace(label string) Logger {
	if len(l.tracer) == 0 {
		l.tracer = label
	} else {
		l.tracer = l.tracer + "." + label
	}

	return l
}
