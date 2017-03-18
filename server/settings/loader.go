package settings

import (
    "os"
    "log"
    "encoding/json"
    "io/ioutil"
    "fmt"
)

const (
    configPath = "settings/"
)

func load(conf *Config) {
    conf.Env = goEnv()
    data, err := ioutil.ReadFile(fmt.Sprintf("%v%v.json", configPath, conf.Env))
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
