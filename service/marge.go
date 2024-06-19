package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab-auto-merge/conf"
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/pkg/hashP"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

const taskFileName = "taskMapInfo.json"

func (p *Service) CreateAutoMergeTask(taskInfo *models.TaskAutoMerge) (err error) {
	if !taskInfo.Check() {
		return fmt.Errorf("taskInfo check fail")
	}
	md5Str := hashP.Md5(taskInfo)
	if len(md5Str) == 0 {
		return fmt.Errorf("md5Str is empty")
	}
	var ctx context.Context
	ctx, taskInfo.Cancel = context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Duration(taskInfo.IntervalTime) * time.Second)
		wg.Done()
		// 初始化任务先执行一次
		if taskInfo.Enable {
			p.creatTask(taskInfo)
		}
		for {
			select {
			case <-ctx.Done():

			case <-ticker.C:
				if !taskInfo.Enable {
					return
				}
				p.creatTask(taskInfo)
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
	taskMapInfo.rwlock.Lock()
	taskMapInfo.taskMap[md5Str] = taskInfo
	taskMapInfo.rwlock.Unlock()

	return
}

func (p *Service) creatTask(taskInfo *models.TaskAutoMerge) {
	var projectMap = make(map[int]string)
	config := conf.GetConfig()
	for _, groupID := range taskInfo.GroupIDs {
		projectInfo, err := p.platform.GetGroupProjects(groupID)
		if err != nil {
			continue
		}
		for _, project := range projectInfo {
			//屏蔽黑名单项目
			if _, ok := config.ProjectBlacklist[project.ID]; ok {
				continue
			}
			if _, ok := projectMap[project.ID]; ok {
				continue
			}
			projectMap[project.ID] = project.Name
			taskInfo.ProjectIDs = append(taskInfo.ProjectIDs, project.ID)
		}
	}
	log.Println("开始检测", projectMap)
	//TODO：由于是使用http 貌似没有必要加上超时
	//cctx, _ := context.WithTimeout(ctx, time.Minute*10)
	for _, projectID := range taskInfo.ProjectIDs {
		req := models.MergeRequest{
			Id:                 projectID,
			SourceBranch:       taskInfo.SourceBranch,
			TargetBranch:       taskInfo.TargetBranch,
			Title:              taskInfo.Title,
			AssigneeId:         p.userID,
			ReviewerIds:        taskInfo.ReviewerID,
			RemoveSourceBranch: taskInfo.RemoveSourceBranch,
		}
		err := p.platform.AutoMerge(req)
		if err != nil {
			log.Println(err)
		}
	}
}

func (p *Service) SaveTaskMapInfo() (err error) {
	_ = os.Remove(taskFileName)

	var info []models.TaskAutoMerge
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

func (p *Service) GetGroups() (groups []models.GroupInfo, err error) {
	return p.platform.GetGroups()
}

func (p *Service) GetProjects() (projects []models.ProjectInfo, err error) {
	return p.platform.GetProjects()
}

func (p *Service) GetGroupProjects(groupID int) (projects []models.ProjectInfo, err error) {
	return p.platform.GetGroupProjects(groupID)
}

func (p *Service) GetUserByName(name string) (users []models.UserInfo, err error) {
	return p.platform.GetUserByName(name)
}

func (p *Service) DelTask() {
	defer log.Println("任务删除完成")
	taskMapInfo.rwlock.Lock()
	for i, i2 := range taskMapInfo.taskMap {
		i2.Cancel()
		delete(taskMapInfo.taskMap, i)
	}
	taskMapInfo.rwlock.Unlock()
}

// LoadTaskMapInfo 通过文件加载任务
func (p *Service) LoadTaskMapInfo() (err error) {
	taskMapByte, err := os.ReadFile(taskFileName)
	if err != nil {
		return
	}
	var taskMap []models.TaskAutoMerge
	err = json.Unmarshal(taskMapByte, &taskMap)
	if err != nil {
		return
	}
	for i, merge := range taskMap {
		if !merge.Check() {
			continue
		}
		err = p.CreateAutoMergeTask(&taskMap[i])
		if err != nil {
			return err
		}
	}
	return
}
