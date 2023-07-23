package main
 
import "fmt"
 
const hello string = "Hello World"
 
func main(){
    fmt.Printf("hello=%s\n", hello)
    const n = 100
    nx2 := n*2
    fmt.Printf("n=%d\n", n)
    fmt.Printf("nx2=%d\n", nx2)
    fmt.Scanf("Pause")
}