package model

import "time"

// FileMeta is used for storing the project information in a
// data.meta.json file. Each project has a unique directory
// that contains a project directory and a data.meta.json file.
type FileMeta struct {
	Name      string    `json:"name"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}
