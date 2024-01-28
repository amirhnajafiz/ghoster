package mongodb

type Config struct {
	Database string `koanf:"database"`
	URI      string `koanf:"uri"`
}
