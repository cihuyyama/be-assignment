package config

type Config struct {
	Srv Server
	DB  Database
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
