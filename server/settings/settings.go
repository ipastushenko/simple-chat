package settings

import (
    "sync"
    "github.com/spf13/viper"
    "github.com/ipastushenko/simple-chat/server/services/websocket"
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
    //TODO: temp init of events
    websocket.NewEventsService().RegisterEvent(
        "message",
        websocket.InitMessageEventHandler,
    )

    conf.SetEnvPrefix(envPrefix)
    conf.AutomaticEnv()
}
