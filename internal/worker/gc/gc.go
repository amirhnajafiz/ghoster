package gc

import (
	"log"
	"time"
)

const (
	functionsDir = "functions"
	prefixToken  = "xxx-"
)

func NewGarbageCollector(interval int) {
	tk := time.NewTicker(time.Duration(interval) * time.Second)

	for {
		<-tk.C

		if count, err := deleteFilesWithStartToken(functionsDir, prefixToken); err != nil {
			log.Printf("gc failed to run: %v\n", err)
		} else if count > 0 {
			log.Printf("gc collected %d items\n", count)
		}
	}
}
