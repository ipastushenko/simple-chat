package settings

import (
    "sync"
    "github.com/spf13/viper"
)

const (
    envPrefix = "SIMPLE_CHAT_SERVER"
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

func initSettings() {
    conf.SetEnvPrefix(envPrefix)
    conf.AutomaticEnv()
}
