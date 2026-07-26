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
		333,
		5,
		100,
		10*time.Second,
		true)

	hydron := hydron.Hydron_Creator(*configs)
	hydron.Run(context.Background())

}
