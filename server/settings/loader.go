package settings

import (
    "os"
    "log"
    "encoding/json"
    "io/ioutil"
)

const (
    configPath = "settings/"
)

func Load(conf *Config) {
    conf.Env = Env()
    data, err := ioutil.ReadFile(configPath + conf.Env + ".json")
    if err != nil {
        log.Println(err.Error())
        os.Exit(1)
    }
    err = json.Unmarshal(data, conf)
    if err != nil {
        log.Println(err.Error())
        os.Exit(1)
    }
}
