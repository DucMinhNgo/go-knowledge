package main

import (
	"fmt"
    "net"
    "net/rpc"
    "time"
    "log"
	"strconv"
    config "server3/config"
)

var server_name = "Server3"
var server_port = "5003"
var current_data = 0
var config_path = "../conf.json"

type DataService struct{
    value string
}

func (p *DataService) Update(request int, reply *string) error {
    current_data = request
    data := strconv.Itoa(request)
    fmt.Println("Received ", data)
    *reply = server_name + " " + data
    // config
    configuration := config.Configuration{}
    configuration = config.Config(config_path);

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

    go rpc.Accept(listener)

    // Call MyFunction every 1 second
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

	// config
	configuration := config.Configuration{}
    configuration = config.Config(config_path);
	PrimaryServer := configuration.Primary

    for range ticker.C {
		// rpc client
		client, err := rpc.Dial("tcp", "localhost:" + PrimaryServer)
		if err != nil {
			log.Println("dead primary server let vote new primary server")
			continue;
		}
		var reply string
		err = client.Call(PrimaryServer + ".CheckAlive", 1, &reply)
		if err != nil {
			// log.Fatal(err)
			// primary server is deadth
		} else {
			
		}
		// fmt.Println(reply)
		client.Close()
    }
}
