package config

import (
	"os"
	"encoding/json"
	"fmt"
    "io/ioutil"
)

type Configuration struct {
    Primary    string
    ServerList   []string
}

// Read Config File
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

// Write Config File
func UpdateConfigFile(path string, data Configuration) bool {
    fmt.Println(path);
    fmt.Println(data);

    file, _ := json.MarshalIndent(data, "", "")

    _ = ioutil.WriteFile(path, file, 0644)

    return true;
}