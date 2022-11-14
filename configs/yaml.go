package configs

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// SimpleBankAccount is a struct for the representing the app's configuration. envconfig tag key has been added to also enable reading properties from envs
type SimpleBankAccount struct {
	Host string `yaml:"host" envconfig:"SIMPLE_BANK_ACCOUNT_HOST"`
	Port string `yaml:"port" envconfig:"SIMPLE_BANK_ACCOUNT_PORT"`
}

// PgDbConfig is a struct for PostgresDb configurations.
type PgDbConfig struct {
	User     string `yaml:"user" envconfig:"POSTGRES_USER"`
	Password string `yaml:"password" envconfig:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host" envconfig:"POSTGRES_HOST"`
	Port     string `yaml:"port" envconfig:"POSTGRES_PORT"`
	DBName   string `yaml:"dbname" envconfig:"POSTGRES_DB"`
}

// YamlConfig maps the configuration in the yaml file into a struct
type YamlConfig struct {
	App        SimpleBankAccount `yaml:"app"`
	PgDatabase PgDbConfig        `yaml:"database"`
}

func ReadYaml(path string) *YamlConfig {
	if path == "" {
		path = defaultYamlConfigPath()
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer func() { _ = f.Close() }()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("error decoding yaml file into config struct: %s\n", err)
		os.Exit(2)
	}
	return &cfg
}

// defaultYamlConfigPath reads the path of current working directory, and then moves a directory up
func defaultYamlConfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("error encountered reading path: %s\n", err)
		os.Exit(2)
	}

	filename := "config.yaml"
	dir := path.Dir(path.Ext(wd))
	dir = path.Join(dir, filename)
	return dir
}
