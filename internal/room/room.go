/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/8/20 3:10 下午
 * @Desc: TODO
 */

package room

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
	ErrNotFoundRoom = errors.New("not found room")
)

type (
	Room struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}

	Manager struct {
		id    int32
		rooms sync.Map
	}
)

func NewManager() *Manager {
	managerOnce.Do(func() {
		managerInstance = &Manager{}
	})
	return managerInstance
}

// AddRoom 添加房间
func (m *Manager) AddRoom(room *Room) int32 {
	room.ID = atomic.AddInt32(&m.id, 1)
	m.rooms.Store(room.ID, room)
	return room.ID
}

// HasRoom 是否存在房间
func (m *Manager) HasRoom(name string) bool {
	_, ok := m.rooms.Load(name)
	return ok
}

// GetRoom 获取房间
func (m *Manager) GetRoom(id int32) (*Room, error) {
	v, ok := m.rooms.Load(id)
	if !ok {
		return nil, ErrNotFoundRoom
	}

	return v.(*Room), nil
}

// DelRoom 删除房间
func (m *Manager) DelRoom(id int32) {
	m.rooms.Delete(id)
}
