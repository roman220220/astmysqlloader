package db

import (
	"database/sql"
	"fmt"
	"astmysqlloader/app/internal/config"
	log "astmysqlloader/app/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"time"
)

type DB struct {
	Config *config.DBConfig
	DBConn *sqlx.DB
	Scheme string
}

func (a *DB) ConnectDB() error {

	cfg := a.Config.GetDBConfig()
	a.Scheme = cfg.Scheme
	var conf string
	if cfg.Driver == "pgx" {
		conf = fmt.Sprintf("host=%s user=%s password=%s database=%s  sslmode=disable",
			cfg.DBServer, cfg.Username, cfg.Password, cfg.DBName)

	}
	if cfg.Driver == "mysql" {

		conf = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&charset=utf8", cfg.Username, cfg.Password, cfg.DBServer, cfg.DBName)

	}

	conn, err := sqlx.Open(cfg.Driver, conf)

	if err != nil {
		log.MakeLog(1, err)
		return err
	}
	conn.SetMaxOpenConns(cfg.MaxOpenConns)
	conn.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Minute)
	conn.SetMaxIdleConns(cfg.MaxIdleConns)
	a.DBConn = conn
	return nil
}
func (a *DB) DBClose() {
	a.DBConn.Close()
}

type Contacts struct {
	Id               int
	AOR              string `db:"aor"`
	Contact          string `db:"contact"`
	Hash             string `db:"hash"`
	Avail            string `db:"avail"`
	Qualify          string `db:"qualify"`
}

func (a *DB) InsertContacts(dialCase map[string]interface{}) {

	sqlStatement := "INSERT INTO " + a.Scheme + ".contacts (aor, contact, hash, qualify) VALUES" +
		"('" + Contacts["aor"].(string) + "','" +
		"" + Contacts["contact"].(string) + "','" +
		"" + Contacts["hash"].(string) + "','" +
		"" + Contacts["avail"].(string) + "','" +
		"" + Contacts["qualify"].(string) + "')"
	_, err := a.DBConn.Exec(sqlStatement)
	if err != nil {
		log.MakeLog(1, err)
	}
	log.MakeLog(3, sqlStatement)
}


