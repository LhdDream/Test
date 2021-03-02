package main


import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

/**
* 功能：rpc server code
* author: 孙松


* email: 741896420@qq.com
 */


type Args struct {
	A,B int
}

type Quotient struct {
	Que,Rem int
}
type Arith int


//乘积
func (t *Arith)Multiply(a *Args,reply *int)  error{
	*reply = a.A * a.B
	return nil
}

func (t *Arith)Divide(a *Args,quo *Quotient) error {
	if a.B == 0{
		return errors.New("divide by zero")
	}
	quo.Que = a.A / a.B
	quo.Rem = a.A % a.B
	return nil
}
func main()  {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l,err := net.Listen("tcp","127.0.0.1:1234")
	if err!= nil{
		log.Fatal("listen error:",err)
	}
	go http.Serve(l,nil)
	fmt.Println("server 启动成功")
	os.Stdin.Read(make([]byte,1))
}