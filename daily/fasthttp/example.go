package main

import "fmt"

func httpHandle(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Printf(ctx, "hello")
}

func main() {
	router := fasthttprouter.New()
	router.Get("/",httpHandle)
	if err := fasthttp.ListenAndServe("0.0.0.0:12345", router.Handler); err != nil {
		fmt.Println("start fasthttp fail:", err.Error())
	}
}
