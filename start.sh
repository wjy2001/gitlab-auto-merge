#!/bin/bash

# 杀死 gitlab-auto-merge 进程
kill -9 $(ps -ef|grep gitlab-auto-merge|grep -v grep|awk '{print $2}')

# 等待一段时间以确保进程完全关闭
sleep 2

# 重新启动 gitlab-auto-merge 服务
nohup ./gitlab-auto-merge &

sleep 1

echo "gitlab-auto-merge 服务已重启"