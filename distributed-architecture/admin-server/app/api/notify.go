package api

import (
	"context"
	"fmt"
	"github.com/dobyte/due-example/internal/pb"
	"github.com/dobyte/due-example/internal/route"
	"github.com/dobyte/due/cluster/master"
	"github.com/dobyte/due/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Notify struct {
	proxy *master.Proxy
}

func NewNotify(proxy *master.Proxy) *Notify {
	return &Notify{proxy: proxy}
}

// Push 发送消息
func (a *Notify) Push(ctx *gin.Context) {
	req := &struct {
		UID     int64  `json:"uid" form:"uid" binding:"required,numeric,min=1"` // 用户ID
		Message string `json:"message" form:"message" binding:"required"`       // 消息
	}{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := a.proxy.Push(context.Background(), req.UID, &master.Message{
		Route: route.NotifyMessage,
		Data:  &pb.Notify{Message: req.Message},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "push success"})
}

// Multicast 组播消息
func (a *Notify) Multicast(ctx *gin.Context) {
	req := &struct {
		UIDS    []int64 `json:"uids" form:"uids" binding:"required"`       // 用户ID
		Message string  `json:"message" form:"message" binding:"required"` // 消息
	}{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	total, err := a.proxy.Multicast(context.Background(), req.UIDS, &master.Message{
		Route: route.NotifyMessage,
		Data:  &pb.Notify{Message: req.Message},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("multicast success, total: %d", total)})
}

// Broadcast 广播消息
func (a *Notify) Broadcast(ctx *gin.Context) {
	req := &struct {
		Message string `json:"message" form:"message" binding:"required"` // 消息
	}{}

	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	total, err := a.proxy.Broadcast(context.Background(), session.User, &master.Message{
		Route: route.NotifyMessage,
		Data:  &pb.Notify{Message: req.Message},
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("broadcast success, total: %d", total)})
}
