package main

import (
	"Tiomon/Pendencias/config"
	"Tiomon/Pendencias/routes"
	"fmt"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	routes.HandleRequests()

}
