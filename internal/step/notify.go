package step

import (
	"context"
	_ "embed"
	"strings"
	"time"

	"github.com/dronestock/drone"
	"github.com/dronestock/feishu/internal/config"
	"github.com/dronestock/feishu/internal/step/internal/constant"
	"github.com/dronestock/feishu/internal/step/internal/feishu/message"
	"github.com/dronestock/feishu/internal/step/internal/notify"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/tpl"
	"github.com/rs/xid"
)

//go:embed internal/notify/template.gohtml
var defaultNotifyTemplate []byte

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

	request := new(notify.Request)
	build := new(notify.Build)
	build.Status = n.base.Value("BUILD_STATUS").String()
	build.Url = n.base.Value("BUILD_LINK").String()
	build.Name = n.base.Value("REPO").String()
	build.Created = n.base.Value("BUILD_CREATED").Time()
	build.Finished = n.base.Value("BUILD_STARTED").Time()
	build.Elapsed = n.base.Elapsed().Truncate(time.Second)
	build.Steps = n.base.Value("FAILED_STEPS").Slices()
	request.Build = build

	code := new(notify.Code)
	code.Pr = n.base.Value("COMMIT_LINK").String()
	code.Repository = n.base.Value("REPO_LINK").String()
	code.Commit = n.base.Value("COMMIT_LINK").String()
	code.Message = n.base.Value("COMMIT_MESSAGE").String()
	request.Code = code

	// 加载模板
	req.Content, err = n.load(request)

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

func (n *Notify) load(req *notify.Request) (content string, err error) {
	if "" != n.card.Template {
		content, err = tpl.New(n.card.Template).File().Data(req).Build().ToString()
	} else {
		content, err = tpl.New(string(defaultNotifyTemplate)).Data(req).Build().ToString()
	}

	if nil == err { // ! 去掉所有空白字符，不然会报格式错误
		content = strings.ReplaceAll(content, "\n", "")
	}

	return
}
