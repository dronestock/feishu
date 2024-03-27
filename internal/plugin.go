package internal

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal/config"
	"github.com/dronestock/feishu/internal/step"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	config.App  `default:"${APP}" json:"app,omitempty"`
	config.Card `default:"${CARD}" json:"card,omitempty"`
	config.User `default:"${USER}" json:"user,omitempty"`

	// 未找到用户时推送的默认用户
	config.Notfound `default:"${NOTFOUND}" json:"notfound,omitempty"`
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewToken(&p.Base, &p.App)).Name("授权").Build(),
		drone.NewStep(step.NewNotify(&p.Base, &p.Card, &p.User, &p.Notfound)).Name("通知").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("app", p.App),
		field.New("card", p.Card),
		field.New("user", p.User),
		field.New("notfound", p.Notfound),
	}
}
