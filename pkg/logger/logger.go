package logger

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Logger interface {
	Error(err error, params ...interface{})
	Info(err error, params ...interface{})
	Debug(err error, params ...interface{})
	Trace(label string) Logger
}

func New(level Level) Logger {
	return &logger{
		level: level,
	}
}

type logger struct {
	tracer string
	level  Level
}

func (l logger) Error(err error, params ...interface{}) {
	if l.level < ErrorLevel {
		return
	}

	labels := make([]string, 0)
	for i := 0; i < len(params); i += 2 {
		labels = append(labels, fmt.Sprintf("{ %s: %s }", params[i], params[i+1]))
	}

	log.Println(fmt.Sprintf("%s [ERROR] %s: %v, [%s]", l.timer(), l.tracer, err, strings.Join(labels, ",")))
}

func (l logger) Info(err error, params ...interface{}) {
	if l.level < InfoLevel {
		return
	}

	labels := make([]string, 0)
	for i := 0; i < len(params); i += 2 {
		labels = append(labels, fmt.Sprintf("{ %s: %s }", params[i], params[i+1]))
	}

	log.Println(fmt.Sprintf("%s [INFO] %s: %v, [%s]", l.timer(), l.tracer, err, strings.Join(labels, ",")))
}

func (l logger) Debug(err error, params ...interface{}) {
	if l.level < DebugLevel {
		return
	}

	labels := make([]string, 0)
	for i := 0; i < len(params); i += 2 {
		labels = append(labels, fmt.Sprintf("{ %s: %s }", params[i], params[i+1]))
	}

	log.Println(fmt.Sprintf("%s [DEBUG] %s: %v, [%s]", l.timer(), l.tracer, err, strings.Join(labels, ",")))
}

func (l logger) Trace(label string) Logger {
	if len(l.tracer) == 0 {
		l.tracer = label
	} else {
		l.tracer = l.tracer + "." + label
	}

	return l
}

func (l logger) timer() string {
	return time.Now().Format(time.DateTime)
}
