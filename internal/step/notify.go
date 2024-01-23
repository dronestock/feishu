package step

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal/config"
	"github.com/dronestock/feishu/internal/step/internal/constant"
	"github.com/dronestock/feishu/internal/step/internal/feishu/message"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/rs/xid"
)

type Notify struct {
	base *drone.Base
	card *config.Card
	user *config.User
}

func NewNotify(base *drone.Base, card *config.Card, user *config.User) *Notify {
	return &Notify{
		base: base,
		card: card,
		user: user,
	}
}

func (n *Notify) Runnable() bool {
	return true
}

func (n *Notify) Run(ctx *context.Context) (err error) {
	if token, ok := (*ctx).Value(constant.ContextKeyToken).(string); !ok {
		err = exception.New().Message("没有正确的授权码").Build()
	} else if request, mre := n.makeRequest(); nil != mre {
		err = mre
	} else {
		err = n.send(ctx, request, token)
	}

	return
}

func (n *Notify) makeRequest() (req *message.Request, err error) {
	req = new(message.Request)
	req.Id = xid.New().String()
	req.Receive = n.user.Id
	req.Type = constant.MessageTypeInteractive

	card := new(message.Card)
	card.Variable = make(map[string]any)
	switch n.base.Value("BUILD_STATUS").String() {
	case constant.DroneStatusSuccess:
		card.Id = n.card.Success
	case constant.DroneStatusFailure:
		card.Id = n.card.Failure
		card.Variable["steps"] = n.base.Value("FAILED_STEPS").String()
	default:
		card.Id = n.card.Success
	}
	card.Variable["name"] = n.base.Value("REPO").String()
	card.Variable["url"] = n.base.Value("BUILD_LINK").String()
	card.Variable["repository"] = n.base.Value("REPO_LINK").String()
	card.Variable["commit"] = n.base.Value("COMMIT_LINK").String()
	card.Variable["message"] = n.base.Value("COMMIT_MESSAGE").String()
	card.Variable["created"] = n.base.Value("BUILD_CREATED").Timestamp()
	card.Variable["finished"] = n.base.Value("BUILD_STARTED").Timestamp()
	card.Variable["elapsed"] = n.base.Elapsed().Truncate(time.Second).String()
	card.Variable["pr"] = n.base.Value("COMMIT_LINK").String()

	content := new(message.Content)
	content.Type = "template"
	content.Data = card
	if bytes, me := json.Marshal(content); nil != me {
		err = me
	} else {
		req.Content = string(bytes)
	}

	return
}

func (n *Notify) send(ctx *context.Context, req *message.Request, token string) (err error) {
	rsp := new(message.Response)
	idType := gox.StringBuilder(n.user.Type, constant.ReceiveTypeStaff).String()
	http := n.base.Http().R().SetContext(*ctx).SetAuthToken(token).SetBody(req).SetResult(rsp)
	if response, pe := http.SetQueryParam(constant.ReceiveType, idType).Post(constant.MessageUrl); nil != pe {
		err = pe
	} else if response.IsError() {
		err = exception.New().Message("飞书返回错误").Field(field.New("response", string(response.Body()))).Build()
	} else {
		n.base.Debug("发送消息成功", field.New("response", rsp))
	}

	return
}
