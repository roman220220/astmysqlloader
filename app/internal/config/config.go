package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/roman220220/astmysqlloader/app/internal/logger"
)

type DBConfig struct {
	Driver          string        `json:"driver"`
	DBServer        string        `json:"db_server"`
	Port            int           `json:"port"`
	DBName          string        `json:"db_name"`
	Scheme          string        `json:"scheme"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
}

func (b *DBConfig) GetDBConfig() *DBConfig {

	f, err := os.Open("./configs/dbconfig.json")
	if err != nil {
		logger.MakeLog(1, err)
	}
	json.NewDecoder(f).Decode(&b)
	return b
}

