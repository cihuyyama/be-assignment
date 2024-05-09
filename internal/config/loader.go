package config

func NewConfig() *Config {
	return &Config{
		Srv: Server{
			Host: "127.0.0.1",
			Port: "5000",
		},
		DB: Database{
			Host:     "127.0.0.1",
			Port:     "5432",
			Name:     "concrete",
			User:     "postgres",
			Password: "postgres",
		},
	}
}
