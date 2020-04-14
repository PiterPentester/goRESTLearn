package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func CreateConfig() *Config {
	return &Config{
		BindAddr: ":3939",
		LogLevel: "debug",
	}
}
