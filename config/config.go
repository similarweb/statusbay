package config

type Metric struct {
	Title  string `yaml:"title"`
	Metric string `yaml:"metric"`
}

type LogConfig struct {
	Level       string `yaml:"level"`
	GelfAddress string `yaml:"gelf_address"`
}
