package main

import "sdkflight/api"

func main() {
	api := api.NewFlightRadar24API()
	api.Login(`admin@tenxtenx.com`, `e9Lc5r7_tCUS!U5`)
}
