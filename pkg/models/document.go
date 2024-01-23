package models

import "time"

type Document struct {
	UUID        string    `bson:"id"`
	Title       string    `bson:"title"`
	User        string    `bson:"user"`
	StoragePath string    `bson:"storage_path"`
	Forbidden   bool      `bson:"forbidden"`
	LastExecute time.Time `bson:"last_execute"`
	CreatedAt   time.Time `bson:"created_at"`
	DeletedAt   time.Time `bson:"deleted_at"`
}
