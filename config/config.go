package config

import "github.com/spf13/viper"

// Config is used to store all the application configuration values
type Config struct {
	logLevel  string
	logFormat string
	docsPath  string
}

// New creates the config from environment or config file
func New() Config {
	return LoadConfig(NewWithViper())
}

func NewWithViper() *viper.Viper {
	vp := viper.New()
	vp.AutomaticEnv()
	vp.SetConfigName("application")
	vp.AddConfigPath("./")
	vp.AddConfigPath("../")
	vp.AddConfigPath("../../")
	vp.ReadInConfig()
	return vp
}

func LoadConfig(vp *viper.Viper) Config {
	return Config{
		logLevel:  getString(vp, "log_level", "debug"),
		logFormat: getString(vp, "log_format", "json"),
		docsPath:  getString(vp, "docs_path"),
	}
}

func (c Config) LogLevel() string {
	return c.logLevel
}

func (c Config) LogFormat() string {
	return c.logFormat
}
