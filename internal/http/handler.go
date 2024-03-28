package http

import "log"

type Handler struct{}

const functionsDir = "functions"

func (h Handler) ListFunctions() {
	functions, err := listDirectoryItems(functionsDir)
	if err != nil {
		return
	}

	log.Println(functions)
}
