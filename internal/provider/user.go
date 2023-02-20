package provider

import (
	"context"
	"github.com/dobyte/due-example/internal/user"
)

type User struct {
}

type ModifyNicknameReq struct {
	Account  string
	Nickname string
}

type ModifyNicknameReply struct {
}

func NewUser() *User {
	return &User{}
}

// ModifyNickname 修改用户昵称
func (p *User) ModifyNickname(ctx context.Context, req *ModifyNicknameReq) (*ModifyNicknameReply, error) {
	u, err := user.NewManager().GetUser(req.Account)
	if err != nil {
		return nil, err
	}

	u.Nickname = req.Nickname

	return &ModifyNicknameReply{}, nil
}
