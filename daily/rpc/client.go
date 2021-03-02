package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A,B int
}

type Quotient struct {
	 Quo,Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp",":1234")
	if err != nil {
		log.Fatal("dialHttp error",err)
		return
	}

	args := &Args{7,8}
	var reply int
	if err = client.Call("Arith.Multiply", args, &reply) ; err != nil {
		log.Fatal("call error ",err)
	}
	fmt.Printf("Multiply %d * %d = %d",args.A,args.B,reply)
	var quo Quotient
	divCall := client.Go("Arith.Divide", args, &quo, nil)
	for {
		select {
		case <-divCall.Done:
			fmt.Printf("%d divide %d是%d %d, 退出执行!", args.A, args.B, quo.Quo,quo.Rem)
			return
		default:
			fmt.Println("继续等待....")
			time.Sleep(time.Second * 1)
		}
	}
}