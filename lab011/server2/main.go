package main

import (
    "log"
    "net"
    "net/rpc"
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "Hellolooooo " + request
    return nil
}

func getConfig () {
    
}

type Data struct {
    primary string
    server_list []string
}

func main() {
    // Open our jsonFile
    content, err := ioutil.ReadFile("config.json")

    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
    
    // defer content.Close()

    var payload Data
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

    fmt.Println(payload.primary)
    // fmt.Println(result["server_list"])

    //
    rpc.RegisterName("HelloService", new(HelloService))
    listener, err := net.Listen("tcp", ":5001")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("Accept error:", err)
        }
        go rpc.ServeConn(conn)
    }
}