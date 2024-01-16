package core

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal/config"
	"github.com/dronestock/feishu/internal/step"
)

type Plugin struct {
	drone.Base
	config.App  `default:"${APP}" json:"app,omitempty"`
	config.Card `default:"${CARD}" json:"card,omitempty"`
	config.User `default:"${USER}" json:"user,omitempty"`
}

func NewPlugin() drone.Plugin {
	return new(Plugin)
}

func (p *Plugin) Config() drone.Config {
	return p
}

func (p *Plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewToken(&p.Base, &p.App)).Name("授权").Build(),
		drone.NewStep(step.NewNotify(&p.Base, &p.Card, &p.User)).Name("通知").Build(),
	}
}
