package platform

//
//import (
//	"fmt"
//	"gitlab-auto-merge/conf"
//	"gitlab-auto-merge/utils"
//)
//
//type Gitea struct {
//	pre *utils.PreRequest
//}
//
//func NewGitea() *Gitea {
//	baseUrl := fmt.Sprint(conf.Config.Parameter.BasicUrl, gitlabBasicAuth)
//	baseHeaders := map[string]string{
//		"Private-Token": conf.Config.Parameter.Token,
//	}
//	return &Gitea{
//		pre: utils.NewPreRequest(utils.InitRequest{
//			BaseURL:         baseUrl,
//			BaseHeaders:     baseHeaders,
//			BaseQueryParams: nil,
//		}),
//	}
//}
//
//// 获取自己的信息
//func (p *Gitea) GetOwnInfo() int {
//}
//
//// 获取用户的ID
//func (p *Gitea) GetUserIDByName(name string) int {
//}
//
//// 获取组下的项目
//func (p *Gitea) GetGroupProjects(groupID string) {
//}
//
//// 获取保护分支
//// GetProtectedBranches()
//// 创建合并请求
//func (p *Gitea) CreateMerge() {
//}
//
//// 获取最后一次提交
//func (p *Gitea) GetLastCommit() {
//}
