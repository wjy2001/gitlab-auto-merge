package service

import (
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/platform"
	"sync"
	"time"
)

func init() {
	taskMapInfo.taskMap = make(map[time.Time]*models.TaskAutoMarge)
}

var taskMapInfo struct {
	taskMap map[time.Time]*models.TaskAutoMarge
	rwlock  sync.RWMutex
}

type Service struct {
	platform platform.Base
}

func NewService(p platform.Base) *Service {
	return &Service{platform: p}
}
