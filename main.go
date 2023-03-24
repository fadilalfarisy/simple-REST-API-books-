package main

import "challenge-1/routers"

func main() {
	var PORT = ":3000"

	routers.StartServer().Run(PORT)
}
