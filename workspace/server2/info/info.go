package info

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type ServerInfo struct {
	ServerName string
	ServerPort string
	ServerVersion string
	CurrentData int
	ConfigPath string
}

func GetServerInfo(path string) ServerInfo {
	file, _ := os.Open(path)
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := ServerInfo{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }

    return configuration
}

func UndateServerInfo(path string, data ServerInfo) bool {
	fmt.Println(path);
    fmt.Println(data);

    file, _ := json.MarshalIndent(data, "", "")

    _ = ioutil.WriteFile(path, file, 0644)

    return true;
}