package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal/core"
)

func main() {
	drone.New(core.NewPlugin).Boot()
}
