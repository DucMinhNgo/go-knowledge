package main

import (
    "log"
    "net"
    "net/rpc"
)

// định nghĩa service struct 
type HelloService struct{}

// định nghĩa hàm service Hello, quy tắc:
// 1. Hàm service phải public (viết hoa)
// 2. Có hai tham số trong hàm
// 3. Tham số thứ hai phải kiểu con trỏ
// 4. Phải trả về kiểu error

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "Hellolooooo " + request
    // trả về error = nil nếu thành công
    return nil
}

func main() {
    rpc.RegisterName("HelloService", new(HelloService))
    // chạy rpc server trên port 1234
    listener, err := net.Listen("tcp", ":1234")
    // nếu có lỗi thì in ra
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }
    // vòng lặp để phục vụ nhiều client
    for {
        // accept một connection đến
        conn, err := listener.Accept()
        // in ra lỗi nếu có
        if err != nil {
            log.Fatal("Accept error:", err)
        }
        // phục vụ client trên một goroutine khác
        // để giải phóng main thread tiếp tục vòng lặp
        go rpc.ServeConn(conn)
    }
}