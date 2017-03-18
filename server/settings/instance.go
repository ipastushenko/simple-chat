package settings

import (
    "sync"
)

var (
    conf *Config
    once sync.Once
)

func GetInstance() *Config {
    once.Do(func() {
        conf = &Config{}
        Load(conf)
    })

    return conf
}
