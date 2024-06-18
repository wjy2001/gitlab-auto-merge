package models

import (
	"context"
	"time"
)

type TaskAutoMerge struct {
	ProjectIDs         []int              `json:"project_ids"`          //项目id
	GroupIDs           []int              `json:"group_ids"`            //分组id
	SourceBranch       string             `json:"source_branch"`        //源分支
	TargetBranch       string             `json:"target_branch"`        //目标分支
	Title              string             `json:"title"`                //标题
	ReviewerID         []int              `json:"reviewer_id"`          //审核人id
	IntervalTime       int                `json:"interval_time"`        //间隔时间
	CreatedTime        time.Time          `json:"created_time"`         //创建时间
	RemoveSourceBranch bool               `json:"remove_source_branch"` //是否删除源分支
	Enable             bool               `json:"enable"`               //是否启用
	Cancel             context.CancelFunc `json:"-"`                    //取消任务
}

func (t TaskAutoMerge) Check() bool {
	if t.IntervalTime < 1 {
		return false
	}
	return true
}
