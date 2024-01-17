package step

import (
	"context"

	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal/config"
	"github.com/dronestock/feishu/internal/step/internal/constant"
	"github.com/dronestock/feishu/internal/step/internal/feishu/token"
	"github.com/goexl/exception"
)

type Token struct {
	base *drone.Base
	app  *config.App
}

func NewToken(base *drone.Base, app *config.App) *Token {
	return &Token{
		base: base,
		app:  app,
	}
}

func (t *Token) Runnable() bool {
	return true
}

func (t *Token) Run(ctx *context.Context) (err error) {
	rsp := new(token.Response)
	req := new(token.Request)
	req.Id = t.app.Id
	req.Secret = t.app.Secret
	if response, pe := t.base.Http().R().SetContext(*ctx).SetBody(req).SetResult(rsp).Post(constant.TokenUrl); nil != pe {
		err = pe
	} else if response.IsError() {
		err = exception.New().Message("飞书返回错误").Build()
	} else {
		*ctx = context.WithValue(*ctx, constant.ContextKeyToken, rsp.Token)
	}

	return
}
