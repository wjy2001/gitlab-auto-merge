package service

import (
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/platform"
	"log"
	"sync"
)

func init() {
	taskMapInfo.taskMap = make(map[string]*models.TaskAutoMerge)
}

var taskMapInfo struct {
	taskMap map[string]*models.TaskAutoMerge
	rwlock  sync.RWMutex
}

type Service struct {
	platform platform.Base
	userID   int
}

func NewService(p platform.Base) *Service {
	user, err := p.GetOwnInfo()
	if err != nil {
		log.Println("获取用户信息失败:", err)
	}
	return &Service{platform: p, userID: user.ID}
}
