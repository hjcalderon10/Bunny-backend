package setting

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type PostgresSettings struct {
	Host               string `envconfig:"PG_HOST" required:"true"`
	Port               string `envconfig:"PG_PORT" required:"true"`
	Name               string `envconfig:"PG_DATABASE" required:"true"`
	Username           string `envconfig:"PG_USER" required:"true"`
	Password           string `envconfig:"PG_PASSWORD" required:"true"`
	DbPQTimeout        int
	MaxConnections     int
	MaxIdleConnections int
	ConnMaxLifeTime    time.Duration
}

var PostgresSetting PostgresSettings

func init() {
	if err := envconfig.Process("", &PostgresSetting); err != nil {
		panic(err.Error())
	}
	PostgresSetting.DbPQTimeout = 30
	PostgresSetting.MaxConnections = 10
	PostgresSetting.MaxIdleConnections = 2
	PostgresSetting.ConnMaxLifeTime = 30 * time.Minute

}
