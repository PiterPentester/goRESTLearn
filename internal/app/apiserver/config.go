package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func CreateConfig() *Config {
	return &Config{
		BindAddr: ":3939",
		LogLevel: "debug",
	}
}
