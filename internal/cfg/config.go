package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() *Config {
	v := viper.New()
	v.SetEnvPrefix("SERV")
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "test")
	v.SetDefault("DBPASS", "test")
	v.SetDefault("DBHOST", "localhost")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "testDB")
	v.AutomaticEnv()

	var cfg Config

	err := v.Unmarshal(&cfg)

	if err != nil {
		log.Panic(err)
	}
	return &cfg
}

func (cfg *Config) GetDBString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
}
