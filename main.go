package main

import (
	"fmt"

	"github.com/convitnhodev/flight-radar-go-sdk/service"
)

func main() {
	api := service.NewFlightRadar24API()
	api.Login(``, ``) // TODO: add user and password to test login
	tmp, _ := api.GetAllFlightWithKey("vn321")
	fmt.Println(service.PrettyPrint(tmp))
	fmt.Printf("Total: %d\n", len(tmp))

}
