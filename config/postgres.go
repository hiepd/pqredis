package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Name        string
	MaxPoolSize int
}

func getPostgresConfig(vp *viper.Viper) PostgresConfig {
	return PostgresConfig{
		Host:        vp.GetString("DB_HOST"),
		Port:        vp.GetInt("DB_PORT"),
		Name:        vp.GetString("DB_NAME"),
		Username:    vp.GetString("DB_USER"),
		Password:    vp.GetString("DB_PASSWORD"),
		MaxPoolSize: vp.GetInt("DB_MAXPOOLSIZE"),
	}
}

func (pc PostgresConfig) String() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", pc.Name, pc.Username, pc.Password, pc.Host, pc.Port)
}
