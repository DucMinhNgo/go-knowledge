package main

import (
    "fmt"
    "log"
    "net/rpc"
	"os"
    "encoding/json"
)

type Configuration struct {
    Primary    string
    ServerList   []string
}

func Config() Configuration {
    file, _ := os.Open("conf.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
    fmt.Println("error:", err)
    }

    return configuration
}

func main() {
	// config
	configuration := Configuration{}
    configuration = Config();
	PrimaryServer := configuration.Primary

	// rpc client
    client, err := rpc.Dial("tcp", "localhost:" + PrimaryServer)
    if err != nil {
        log.Fatal("dialing:", err)
    }
    var reply string
    err = client.Call(PrimaryServer + ".Update", 20, &reply)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(reply)
}