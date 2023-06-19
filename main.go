package main

import (
	"fmt"
	api_ "github.com/convitnhodev/flight-radar-go-sdk/service"
)

func main() {
	api := api_.NewFlightRadar24API()
	api.Login(`admin@tenxtenx.com `, `e9Lc5r7_tCUS!U5`)
	tmp, _ := api.GetAllFlightWithKey("vn321")
	fmt.Println(tmp)

}
