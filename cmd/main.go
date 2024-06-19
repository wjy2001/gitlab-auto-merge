package main

import (
	"gitlab-auto-merge/conf"
	"gitlab-auto-merge/platform"
	"gitlab-auto-merge/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 无图形版本
func main() {
	conf.InitConfig()
	s := service.NewService(platform.NewGitlab())
	err := s.LoadTaskMapInfo()
	if err != nil {
		log.Println("加载任务失败：", err)
		s.DelTask()
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		log.Println("退出")
		s.DelTask()
		os.Exit(0)
	}

}
