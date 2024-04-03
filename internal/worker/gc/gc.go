package gc

import (
	"log"
	"time"
)

func NewGarbageCollector(functionsDir, prefixToken string, interval int) {
	go func() {
		tk := time.NewTicker(time.Duration(interval) * time.Second)

		for {
			<-tk.C

			if count, err := deleteFilesWithStartToken(functionsDir, prefixToken); err != nil {
				log.Printf("gc failed to run: %v\n", err)
			} else if count > 0 {
				log.Printf("gc collected %d items\n", count)
			}
		}
	}()
}
