package main

import "github.com/dpcamargo/fullcycle-api/configs"

func main() {
	_, err := configs.LoadConfig("./cmd/server")
	if err != nil {
		panic(err)
	}
}
