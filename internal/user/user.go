package user

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	managerOnce     sync.Once
	managerInstance *Manager
)

var (
	ErrNotFoundUser = errors.New("not found user")
)

type (
	User struct {
		ID       int32  `json:"id"`
		Nickname string `json:"nickname"`
		Account  string `json:"account"`
		Password string `json:"password"`
		Age      int    `json:"age"`
	}

	Manager struct {
		id    int32
		users sync.Map
	}
)

func NewManager() *Manager {
	managerOnce.Do(func() {
		managerInstance = &Manager{}
	})
	return managerInstance
}

// AddUser 添加用户
func (m *Manager) AddUser(user *User) int32 {
	user.ID = atomic.AddInt32(&m.id, 1)
	m.users.Store(user.Account, user)
	return user.ID
}

// HasUser 是否存在用户
func (m *Manager) HasUser(account string) bool {
	_, ok := m.users.Load(account)
	return ok
}

// GetUser 获取用户
func (m *Manager) GetUser(account string) (*User, error) {
	v, ok := m.users.Load(account)
	if !ok {
		return nil, ErrNotFoundUser
	}

	return v.(*User), nil
}
