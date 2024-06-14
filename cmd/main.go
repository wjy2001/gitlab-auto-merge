package main

import (
	"gitlab-auto-merge/platform"
	"gitlab-auto-merge/service"
)

// 无图形版本
func main() {
	s := service.NewService(platform.NewGitlab())
	err := s.LoadTaskMapInfo()
	if err != nil {
		panic(err)
	}
	defer s.DelTask()
	select {}
}
