package main

import (
	api "github.com/lifeng342/go-demo/kitex_gen/api/hello"
	"log"
)

func main() {
	//Replace()
	ChatComplete()
	svr := api.NewServer(new(HelloImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
