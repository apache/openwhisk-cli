package common

import (
    "github.com/spf13/viper"
    "io/ioutil"
    "os"
)

type Wskprops struct {
    APIHost string
    APIVersion string
    AuthKey string
    Namespace string
    ControllerHost string
    ControllerPort string
}

func GetWskprops() *Wskprops {
    var dep Wskprops
    dep.APIHost = ""
    dep.AuthKey = ""
    dep.APIVersion = "v1"
    dep.Namespace = "_"

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
        dep.ControllerHost = viper.GetString("router.host")
        dep.ControllerPort = viper.GetString("controller.host.port")
    }
    return &dep
}
