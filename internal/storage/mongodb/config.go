package mongodb

type Config struct {
	Collection string `koanf:"collection"`
	Database   string `koanf:"database"`
	URI        string `koanf:"uri"`
}
