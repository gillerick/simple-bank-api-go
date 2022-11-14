package configs

import "fmt"

type App struct {
	Host string
	Port string
}

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

// String makes it possible to extend the properties of existing db configuration
func (d Database) String(sslmode string) string {
	return fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"dbname=%s "+
		"password=%s "+
		"sslmode=%s"+
		"", d.Host, d.Port, d.User, d.DBName, d.Password, sslmode)
}

// Config struct maps the yaml config into a Go struct
type Config struct {
	App App
	DB  Database
}

func GetConfig(cfg YamlConfig) Config {
	return Config{App: App{
		Host: cfg.App.Host,
		Port: cfg.App.Port,
	}, DB: Database{
		User:     cfg.PgDatabase.User,
		Password: cfg.PgDatabase.Password,
		Host:     cfg.PgDatabase.Host,
		Port:     cfg.PgDatabase.Port,
		DBName:   cfg.PgDatabase.DBName,
	}}
}
