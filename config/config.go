package config

import (
    "fmt"
    "strings"

    "github.com/spf13/viper"
)

type Configuration struct {
    core    *viper.Viper
}

func (configuration *Configuration) GetBool(key string) bool {
    return configuration.core.GetBool(key)
}

func (configuration *Configuration) GetInt(key string) int {
    return configuration.core.GetInt(key)
}

func (configuration *Configuration) GetFloat64(key string) float64 {
    return configuration.core.GetFloat64(key)
}

func (configuration *Configuration) GetString(key string) string {
    return configuration.core.GetString(key)
}

func New(name string) *Configuration {
    var file string

    if strings.Contains(name, "/") {
        file = name
    } else {
        file = fmt.Sprintf("./config/%s.yaml", name)
    }

    core := viper.New()
    core.SetConfigFile(file)
    if err := core.ReadInConfig(); err != nil {
        panic(err)
    }

    return &Configuration{core: core}
}
