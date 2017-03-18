package main

import (
    "fmt"
    "github.com/ipastushenko/simple-chat/settings"
)

func main () {
    config := settings.GetInstance()
    fmt.Println(config.Env)
    fmt.Println(config.Server.Port)
}
