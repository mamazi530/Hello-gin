package main

import "Hello-gin/router"

func main() {
	// 装载路由
	r := router.NewRouter()
	r.Run(":8080")

}