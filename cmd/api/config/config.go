package config

type Config struct {
	Service Service    `json:"service"`
	Redis   DataSource `yaml:"redis"`
	MariaDB DataSource `yaml:"mariadb"`
}

type Service struct {
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type DataSource struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func Get() *Config {
	return nil
}
