package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal"
)

func main() {
	drone.New(internal.New).Boot()
}
