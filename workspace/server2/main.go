package main

import (
	"fmt"
    "net"
    "net/rpc"
    "time"
    "log"
    config "server2/config"
    serverInfo "server2/info"
    service "server2/service"
)

var server_name = ""
var server_port = ""
var current_data = 0
var config_path = ""

func main() {
    server := serverInfo.ServerInfo{}
    server = serverInfo.GetServerInfo("./info.json")
    server_name = server.ServerName
    server_port = server.ServerPort
    current_data = server.CurrentData

	fmt.Println("Server " + server_name + " is running.....")
	// rpc server
    rpc.RegisterName(server_port, new(service.DataService))
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
    configuration = config.GetConfigFile(config_path)
	PrimaryServer := configuration.Primary

    for range ticker.C {
		// rpc client
		client, err := rpc.Dial("tcp", "localhost:" + PrimaryServer)
        var clientService = client;
		if err != nil {
            configuration.Primary = "5003"
			log.Println("dead primary server let vote new primary server")
            config.UpdateConfigFile(config_path, configuration);
            client, _ := rpc.Dial("tcp", "localhost:" + configuration.Primary)
            clientService = client
			continue;
		}
		var reply string
		err = clientService.Call(PrimaryServer + ".CheckAlive", 1, &reply)
		if err != nil {
			// log.Fatal(err)
			// primary server is deadth
		} else {
			
		}
		// fmt.Println(reply)
		clientService.Close()
    }
}
