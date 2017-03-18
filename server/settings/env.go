package settings

import "os"

const (
  defaultEnv string = "development"
  envName string = "GO_ENV"
)

func Env() string {
    env, ok := os.LookupEnv(envName)

    if ok {
        return env
    }

    return defaultEnv
}
