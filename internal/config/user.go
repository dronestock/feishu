package config

type User struct {
	// 接收用户标识
	Id string `default:"${DRONE_COMMIT_AUTHOR=${USERID}}" json:"id,omitempty"`
	// 接收用户类型
	Type string `default:"${USERTYPE=user}" json:"type,omitempty" validate:"oneof=open user union email chat"`
}

func (u *User) Userid() string {
	return u.Id
}

func (u *User) Usertype() string {
	return u.Type
}
