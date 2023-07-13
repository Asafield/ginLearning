package main

import "ginLearning/routers"

func main() {
	router := routers.InitRouter()
	router.Run()
}
