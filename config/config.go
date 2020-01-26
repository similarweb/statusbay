package config

// MySQLConfig client configuration
type MySQLConfig struct {
	DNS      string `yaml:"dns"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

// SlackConfig configuration
type SlackConfig struct {
	Token           string   `yaml:"token"`
	DefaultChannels []string `yaml:"default_channels"`
}
