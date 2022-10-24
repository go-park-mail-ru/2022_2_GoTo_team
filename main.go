package main

import (
	"2022_2_GoTo_team/server"
)

const serverAddress = ":8080"

var allowOriginsAddressesCORS = []string{"http://localhost:8080"}

func main() {
	server.Run(serverAddress, allowOriginsAddressesCORS)
}