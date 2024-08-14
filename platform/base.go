package platform

import "gitlab-auto-merge/models"

type Base interface {
	// NewPre 初始化请求
	NewPre()
	// GetOwnInfo 获取自己的信息
	GetOwnInfo() (user models.UserInfo, err error)
	// GetUserByName 获取用户的信息
	GetUserByName(name string) (users []models.UserInfo, err error)
	// GetGroups 获取用户的群组
	GetGroups() (groups []models.GroupInfo, err error)
	// GetGroupProjects 获取组下的项目
	GetGroupProjects(groupID int) (projects []models.ProjectInfo, err error)
	// GetProjects 获取能够查看的项目
	GetProjects() (projects []models.ProjectInfo, err error)
	// CreateMerge 创建合并请求
	CreateMerge(body models.MergeRequest) (err error)
	// GetBranch 获取分支详情
	GetBranch(projectID int, branchName string) (branch models.BranchInfo, err error)
	// GetProjectBranches 获取项目下的分支
	GetProjectBranches(projectID int) (branches []models.BranchInfo, err error)
	// GetCommitBranches 获取提交存在的分支
	GetCommitBranches(projectID int, sha string) (branches []string, err error)
	// AutoMerge 自动提交合并请求
	AutoMerge(req models.MergeRequest) (err error)
}
