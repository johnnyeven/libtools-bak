package mysql

import (
	"fmt"
	"net/url"
	"time"

	"golib/tools/conf"
	"golib/tools/conf/presets"
	"golib/tools/sqlx"
)

type MySQL struct {
	Name            string
	Host            string `conf:"upstream"`
	Port            int
	User            string           `conf:"env"`
	Password        presets.Password `conf:"env"`
	Extra           string
	PoolSize        int
	ConnMaxLifetime time.Duration
	presets.Retry
	db *sqlx.DB
}

func (m MySQL) DockerDefaults() conf.DockerDefaults {
	return conf.DockerDefaults{
		"Host": conf.RancherInternal("tool-dbs", m.Name),
		"Port": 3306,
	}
}

func (m MySQL) MarshalDefaults(v interface{}) {
	if mysql, ok := v.(*MySQL); ok {
		mysql.Retry.MarshalDefaults(&mysql.Retry)

		if mysql.Port == 0 {
			mysql.Port = 3306
		}

		if mysql.PoolSize == 0 {
			mysql.PoolSize = 10
		}

		if mysql.ConnMaxLifetime == 0 {
			mysql.ConnMaxLifetime = 4 * time.Hour
		}

		if mysql.Extra == "" {
			values := url.Values{}
			values.Set("charset", "utf8")
			values.Set("parseTime", "true")
			values.Set("interpolateParams", "true")
			values.Set("autocommit", "true")
			values.Set("loc", "Local")
			mysql.Extra = values.Encode()
		}
	}
}

func (m MySQL) GetConnect() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/?%s", m.User, m.Password, m.Host, m.Port, m.Extra)
}

func (m *MySQL) Connect() error {
	m.MarshalDefaults(m)
	db, err := sqlx.Open("logger:mysql", m.GetConnect())
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(m.PoolSize)
	db.SetMaxIdleConns(m.PoolSize / 2)
	db.SetConnMaxLifetime(m.ConnMaxLifetime)
	m.db = db
	return nil
}

func (m *MySQL) Init() {
	if m.db == nil {
		m.Do(m.Connect)
	}
}

func (m *MySQL) Get() *sqlx.DB {
	if m.db == nil {
		panic(fmt.Errorf("get db before init"))
	}
	return m.db
}

type DBGetter interface {
	Get() *sqlx.DB
}
