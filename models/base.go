package models

import (
	"time"
)

type UserInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type GroupInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
}

type ProjectInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type CommitInfo struct {
	ID            string `json:"id"` // 提交id(md5)
	Title         string `json:"title"`
	Message       string `json:"message"`
	AuthorEmail   string `json:"author_email"`
	AuthorName    string `json:"author_name"`
	AuthoredDate  string `json:"authored_date"`
	CommittedDate string `json:"committed_date"`
}

type BranchInfo struct {
	Name   string     `json:"name"`
	Commit CommitInfo `json:"commit"`
}

type MergeRequest struct {
	Id                 int    `json:"id"`                   //项目id
	SourceBranch       string `json:"source_branch"`        //源分支
	TargetBranch       string `json:"target_branch"`        //目标分支
	Title              string `json:"title"`                //标题
	AssigneeId         int    `json:"assignee_id"`          //指派人id
	ReviewerIds        []int  `json:"reviewer_ids"`         //审核人id
	RemoveSourceBranch bool   `json:"remove_source_branch"` //是否删除源分支
}

type MergeInfo struct {
	Id           int         `json:"id"`            //合并请求id
	Iid          int         `json:"iid"`           //合并请求id
	ProjectId    int         `json:"project_id"`    //项目id
	Title        string      `json:"title"`         //标题
	Description  interface{} `json:"description"`   //描述
	State        string      `json:"state"`         //状态
	TargetBranch string      `json:"target_branch"` //目标分支
	SourceBranch string      `json:"source_branch"` //源分支
	Author       UserInfo    `json:"author"`        //创建人
	Assignees    []UserInfo  `json:"assignees"`     //指派人
	Reviewers    []UserInfo  `json:"reviewers"`     //审核人
	CreatedAt    time.Time   `json:"created_at"`    //创建时间
	UpdatedAt    time.Time   `json:"updated_at"`    //更新时间
}

type MergeSimpleInfo struct {
	Id          int         `json:"id"`          //合并请求id
	Iid         int         `json:"iid"`         //合并请求id
	ProjectId   int         `json:"project_id"`  //项目id
	Title       string      `json:"title"`       //标题
	Description interface{} `json:"description"` //描述
	State       string      `json:"state"`       //状态
	CreatedAt   time.Time   `json:"created_at"`  //创建时间
	UpdatedAt   time.Time   `json:"updated_at"`  //更新时间
	WebUrl      string      `json:"web_url"`     //链接
}

type GetMergeReq struct {
	Search       string `json:"search"`
	State        string `json:"state"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
}

func (s *GetMergeReq) ToStringMap() map[string]string {
	res := map[string]string{
		"search":        s.Search,
		"state":         s.State,
		"source_branch": s.SourceBranch,
		"target_branch": s.TargetBranch,
	}
	return res
}

type UpdateMergeReq struct {
	Id          int         `json:"id"`          //合并请求id
	Iid         int         `json:"iid"`         //合并请求id
	ProjectId   int         `json:"project_id"`  //项目id
	Title       string      `json:"title"`       //标题
	Description interface{} `json:"description"` //描述
	State       string      `json:"state"`       //状态
}
