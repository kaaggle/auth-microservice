package core

type Config struct {
	BaseURL string
	Secret  string
	Database
}

type Database struct {
	URL string
}

func NewConfig() *Config {
	return &Config{
		BaseURL: "localhost:5000",
		// Database: Database{"127.0.0.1"},
		Database: Database{"mongodb://school-system:school-system1@ds237192.mlab.com:37192/school-system"},
		// Database: Database{os.Getenv("KAAGLE_DB_URL")},
		Secret: "MySECRET",
	}
}
