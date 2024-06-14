package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab-auto-merge/models"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

const taskFileName = "taskMapInfo.json"

func (p *Service) CreateAutoMargeTask(taskInfo *models.TaskAutoMarge) (err error) {
	if !taskInfo.Check() {
		return fmt.Errorf("taskInfo check fail")
	}

	var ctx context.Context
	ctx, taskInfo.Cancel = context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Duration(taskInfo.IntervalTime) * time.Second)
		wg.Done()
		// 初始化任务先执行一次
		p.creatTask(*taskInfo)
		for {
			select {
			case <-ctx.Done():

			case <-ticker.C:
				p.creatTask(*taskInfo)
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
	taskMapInfo.rwlock.RLock()
	taskMapInfo.taskMap[taskInfo.CreatedTime] = taskInfo
	taskMapInfo.rwlock.RUnlock()

	err = saveTaskMapInfo()
	if err != nil {
		return
	}
	return
}

func (p *Service) creatTask(taskInfo models.TaskAutoMarge) {
	fmt.Println("开始检测")
	for _, groupID := range taskInfo.GroupIDs {
		projectInfo, err := p.platform.GetGroupProjects(groupID)
		if err != nil {
			continue
		}
		for _, project := range projectInfo {
			taskInfo.ProjectIDs = append(taskInfo.ProjectIDs, project.ID)
		}
	}
	//TODO：由于是使用http 貌似没有必要加上超时
	//cctx, _ := context.WithTimeout(ctx, time.Minute*10)
	for _, projectID := range taskInfo.ProjectIDs {
		req := models.MergeRequest{
			Id:                 projectID,
			SourceBranch:       taskInfo.SourceBranch,
			TargetBranch:       taskInfo.TargetBranch,
			Title:              taskInfo.Title,
			AssigneeId:         0,
			ReviewerIds:        taskInfo.ReviewerID,
			RemoveSourceBranch: false,
		}
		err := p.platform.AutoMarge(req)
		if err != nil {
			log.Println(err)
		}
	}
}

func saveTaskMapInfo() (err error) {
	_ = os.Remove(taskFileName)

	var info []models.TaskAutoMarge
	taskMapInfo.rwlock.Lock()
	for _, v := range taskMapInfo.taskMap {
		info = append(info, *v)
	}
	taskMapInfo.rwlock.Unlock()

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		err = fmt.Errorf("序列化失败:%s", err.Error())
		return
	}
	err = os.WriteFile(taskFileName, jsonInfo, 0666)
	if err != nil {
		err = fmt.Errorf("写入文件失败:%s", err.Error())
		return
	}
	return
}

func (p *Service) GetGroups() (groups []*models.GroupInfo, err error) {
	return p.platform.GetGroups()
}

func (p *Service) GetProjects() (projects []*models.ProjectInfo, err error) {
	return p.platform.GetProjects()
}

func (p *Service) GetGroupProjects(groupID int) (projects []*models.ProjectInfo, err error) {
	return p.platform.GetGroupProjects(groupID)
}

func (p *Service) GetUserByName(name string) (users []*models.UserInfo, err error) {
	return p.platform.GetUserByName(name)
}

func (p *Service) DelTask() {
	taskMapInfo.rwlock.RLock()
	for i, i2 := range taskMapInfo.taskMap {
		i2.Cancel()
		delete(taskMapInfo.taskMap, i)
	}
	taskMapInfo.rwlock.RUnlock()
}

// 通过文件加载任务
func (p *Service) LoadTaskMapInfo() (err error) {
	taskMapByte, err := os.ReadFile(taskFileName)
	if err != nil {
		return
	}
	var taskMap []models.TaskAutoMarge
	err = json.Unmarshal(taskMapByte, &taskMap)
	if err != nil {
		return
	}
	for _, marge := range taskMap {
		if !marge.Check() || !marge.Enable {
			continue
		}
		err = p.CreateAutoMargeTask(&marge)
		if err != nil {
			return err
		}
	}
	return
}
