package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
    Primary    string
    ServerList   []string
}

func Config(path string) Configuration {
    file, _ := os.Open(path)
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }

    return configuration
}