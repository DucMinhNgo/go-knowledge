package service

import (
	"fmt"
	"strconv"
	config "server2/config"
	"net/rpc"
)

var server_name = ""
var server_port = ""
var current_data = 0
var config_path = ""

type DataService struct{
    value string
}

func (p *DataService) Update(request int, reply *string) error {
    // service
    current_data = request
    data := strconv.Itoa(request)
    fmt.Println("Received ", data)
    *reply = server_name + " " + data
    // config
    configuration := config.Configuration{}
    configuration = config.GetConfigFile(config_path)

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