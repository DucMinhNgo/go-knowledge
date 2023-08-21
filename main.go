package main

import (
	"fmt" 
	"os"
	"time"
	"sync"
)

var count_on = 0
var count_off = 0
var m sync.Mutex


func get_sum(i int) int {
	s := 0
    for i > 0{
		s = s + i
        i--;
    }
	return s;
}

func create_file() {
	f, err := os.Create("data.txt")

    if err != nil {
        // log.Fatal(err)
    }

	// Chạy sau khi hàm kết thúc
    defer f.Close()

    _, err2 := f.WriteString("Duc Ngo\n")

    if err2 != nil {
        // log.Fatal(err2)
    }

    fmt.Println("done")
}

func mapTry() {
	mySet := make(map[string]string)

	mySet["key1"] = "apple"
	mySet["key2"] = "banana"
	mySet["key3"] = "arantino"

	for key := range mySet {
		fmt.Printf("key=%s value=%s\n", key, mySet[key])
	}
}

func arrayTry() {
	arr := [10]string{"S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i]);
	}
}

func (s Server) PrintServer() {
	fmt.Println("print", s.name)
	fmt.Println("print", s.counter)
	fmt.Println("print", s.leader)
}

type Server struct {
	name string
	counter int
	leader bool
}

func hello(i int) {
	
	fmt.Printf("Hello, world! %d \n", i)
	if i % 2 == 1 {
		count_on += 1
	} else {
		count_off += 1
	}
}

var channel = make(chan int)

func setChannel(i int) {
	if i % 2 == 1 {
		channel <- i
	} else {
		a :=  <- channel
		fmt.Println(a)
	}
}

func getTimer() {
	for {
        fmt.Println("https://gosamples.dev is the best")
        time.Sleep(1 * time.Second)
    }
}


func main () {
	// server1 := Server{
	// 	name: "server1",
	// 	counter: 1,
	// 	leader: true,
	// }

	// server2 := Server{
	// 	name: "server2",
	// 	counter: 2,
	// 	leader: false,
	// }
	
	// array:= [2]Server{server1, server2}
	
	// for _,val := range array {
	// 	fmt.Println(val.name)
	// 	fmt.Println(val.counter)
	// 	fmt.Println(val.leader)
	// }
	
	// server1.PrintServer()
	// m.Lock()
	// for i:=0; i<100; i++ {
	// 	go hello(i)
	// }
	// m.Unlock()

	// // time.Sleep(10)
	// fmt.Println(count_on)
	// fmt.Println(count_off)

	// go setChannel(1)
	// for i:=0; i<1000000; i++ {
	// 	// go setChannel(i)
	// 	fmt.Println("Tick")
	// 	time.Sleep(1)
	// }

	go getTimer()

	time.Sleep(10)
	
	fmt.Println("Hello, world!")
}