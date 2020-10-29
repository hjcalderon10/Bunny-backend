package setting

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
)

var Commons commons

type commons struct {
	AppEnv string `envconfig:"APP_ENV" default:"development"`

	Port string `envconfig:"PORT" default:"8080"`
	Host string `envconfig:"HOST" default:"0.0.0.0"`

	XApplicationID string `envconfig:"APPLICATION_ID" default:"bunny-backend/0.0.1"`
	ProjectName    string
	ProjectVersion string
}

func init() {
	if err := envconfig.Process("", &Commons); err != nil {
		panic(err.Error())
	}

	infoApp := strings.Split(Commons.XApplicationID, "/")
	Commons.ProjectName = infoApp[0]
	Commons.ProjectVersion = strings.Join(infoApp[1:], "")
}
