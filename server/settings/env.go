package settings

import "os"

const (
  defaultEnv string = "development"
  envName string = "GO_ENV"
)

func goEnv() string {
    env, ok := os.LookupEnv(envName)

    if ok {
        return env
    }

    return defaultEnv
}
