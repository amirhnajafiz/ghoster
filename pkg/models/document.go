package models

import "time"

type Document struct {
	UUID        string    `bson:"uuid"`
	Title       string    `bson:"title"`
	StoragePath string    `bson:"storage_path"`
	Forbidden   bool      `bson:"forbidden"`
	LastExecute time.Time `bson:"last_execute"`
	CreatedAt   time.Time `bson:"created_at"`
}
