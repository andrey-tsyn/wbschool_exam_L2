package configuration

type Config struct {
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	Port        string `env:"PORT" envDefault:"8080"`
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`
}
