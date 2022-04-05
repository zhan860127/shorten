package main

import (
	"fmt"

	_ "shorten/controllers"
	"shorten/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

	f := middlewares.dump_to_python('a')
	fmt.Println(f)
}
