package platform

import "gitlab-auto-merge/models"

type Base interface {
	// 获取自己的信息
	GetOwnInfo() (user *models.UserInfo, err error)
	// 获取用户的信息
	GetUserByName(name string) (users []*models.UserInfo, err error)
	// 获取用户的群组
	GetGroups() (groups []*models.GroupInfo, err error)
	// 获取组下的项目
	GetGroupProjects(groupID int) (projects []*models.ProjectInfo, err error)
	// 获取能够查看的项目
	GetProjects() (projects []*models.ProjectInfo, err error)
	// 创建合并请求
	CreateMerge(body models.MergeRequest) (err error)
	// 获取分支详情
	GetBranch(projectID int, branchName string) (branch *models.BranchInfo, err error)
	// 获取提交存在的分支
	GetCommitBranches(projectID int, sha string) (branches []string, err error)
	// 自动提交合并请求
	AutoMarge(req models.MergeRequest) (err error)
}
