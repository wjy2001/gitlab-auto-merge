package main

import (
	"fmt"
	"gitlab-auto-merge/conf"
	"gitlab-auto-merge/platform"
	"gitlab-auto-merge/service"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

// 无图形版本
func main() {
	f, err := os.Create("./profile.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if err = pprof.StartCPUProfile(f); err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()
	if err = pprof.WriteHeapProfile(f); err != nil {
		fmt.Println(err)
		return
	}

	conf.InitConfig()
	s := service.NewService(platform.NewGitlab())
	err = s.LoadTaskMapInfo()
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
