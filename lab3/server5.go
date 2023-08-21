package main

import (
    "fmt"
    "os"
    "encoding/json"
    "log"
    "net"
    "net/rpc"
    "strconv"
)

var server_name = "Server5"
var server_port = "5005"
var current_data = 0

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

type DataService struct{
    value string
}

func (p *DataService) Update(request int, reply *string) error {
    current_data = request
    data := strconv.Itoa(request)
    fmt.Println("Received ", data)
    *reply = server_name + " " + data
    // config
    configuration := Configuration{}
    configuration = Config();

    for _, data := range configuration.ServerList {
        if data != server_port {
            OtherServer := data
            // rpc sync to other server
            // rpc client
            client, err := rpc.Dial("tcp", "localhost:" + OtherServer)
            if err != nil {
                // log.Fatal("dialing:", err)
            } else {
                var reply string
                err = client.Call(OtherServer + ".Sync", request, &reply)
                if err != nil {
                    // log.Fatal(err)
                }

                fmt.Println("Sync", request, "to", data)
                fmt.Println(reply)
            }
        } 
    }
 
    return nil
}

func (p *DataService) Sync(request int, reply *string) error {
    // config
    data := strconv.Itoa(request)
    fmt.Println("Received ", data)
    *reply = server_name + " " + data
    return nil
}

func main() {
    fmt.Println("Server " + server_name + " is running.....")
    // rpc server
    rpc.RegisterName(server_port, new(DataService))
    listener, err := net.Listen("tcp", ":" + server_port)
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

