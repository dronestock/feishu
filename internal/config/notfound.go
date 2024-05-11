package config

type Notfound struct {
	// 接收用户标识
	Id string `default:"${DRONE_REPO_OWNER}" json:"id,omitempty"`
	// 接收用户类型
	Type string `default:"user" json:"type,omitempty" validate:"oneof=open user union email chat"`
}

func (n *Notfound) Userid() string {
	return n.Id
}

func (n *Notfound) Usertype() string {
	return n.Type
}
