package main

import (
	"fmt"

	"shorten/routers"
)

func main() {
	r := routers.SetupRouter()

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
