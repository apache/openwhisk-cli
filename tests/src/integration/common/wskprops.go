package common

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

type Wskprops struct {
	APIHost string
	AuthKey string
}

func GetWskprops() *Wskprops {
	var dep Wskprops
	dep.APIHost = ""
	dep.AuthKey = ""

	viper.SetConfigName("whisk")
	viper.AddConfigPath(os.Getenv("OPENWHISK_HOME"))

	err := viper.ReadInConfig()
	if err == nil {
		authPath := viper.GetString("testing.auth")

		b, err := ioutil.ReadFile(authPath)
		if err == nil {
			dep.AuthKey = string(b)
		}
		dep.APIHost = viper.GetString("router.host")
	}
	return &dep
}
