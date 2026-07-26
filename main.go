package main

import (
	"context"
	"time"

	"main/hydron"
)

func main() {

	configs := hydron.Create_packetjob(
		"127.0.0.1",
		"127.0.0.1",
		8080,
		"/",
		5,
		10,
		3*time.Second,
		true)

	hydron := hydron.Hydron_Creator(*configs)
	hydron.Run(context.Background())

}
