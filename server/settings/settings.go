package settings

import (
    "sync"
    "os"
    "log"
    "fmt"
    "github.com/spf13/viper"
)

const (
    defaultEnv string = "development"
    envName string = "GO_ENV"
    configPath = "./settings"
    configType = "json"
)

var (
    conf *viper.Viper
    once sync.Once
)

func GetInstance() *viper.Viper {
    once.Do(func() {
        conf = viper.New()
        initSettings()
    })

    return conf
}

func goEnv() string {
    if env, ok := os.LookupEnv(envName); ok {
        return env
    }

    return defaultEnv
}

func initSettings() {
    conf.Set("Env", goEnv())
    conf.AddConfigPath(configPath)
    conf.SetConfigType(configType)
    conf.SetConfigName(fmt.Sprintf("config.%v", conf.GetString("Env")))
    err := conf.ReadInConfig()
    if err != nil {
        log.Println(err.Error())
        os.Exit(1)
    }
}
