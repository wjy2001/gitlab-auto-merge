package platform

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab-auto-merge/conf"
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/pkg/httpP"
	"net/http"
	"strconv"
	"strings"
)

const (
	gitlabBasicAuth = "/api/v4"
)

type Gitlab struct {
	pre *resty.Client
}

func NewGitlab() *Gitlab {
	c := conf.GetConfig()
	baseUrl := fmt.Sprint(c.Parameter.BasicUrl, gitlabBasicAuth)
	baseHeaders := map[string]string{
		"Private-Token": c.Parameter.Token,
	}

	return &Gitlab{
		pre: httpP.NewPreRequestClient(httpP.InitRequest{
			BaseURL:         baseUrl,
			BaseHeaders:     baseHeaders,
			BaseQueryParams: nil,
		}),
	}
}

// 获取自己的信息
func (p *Gitlab) GetOwnInfo() (user *models.UserInfo, err error) {
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Method:      http.MethodGet,
		Url:         "/user",
		Headers:     nil,
		QueryParams: nil,
		Body:        nil,
	})

	res, err := pre.GetRespBody()
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &user)
	if err != nil {
		return
	}
	return
}

// 获取用户的信息
func (p *Gitlab) GetUserByName(name string) (users []*models.UserInfo, err error) {
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Method:  http.MethodGet,
		Url:     "/users",
		Headers: nil,
		QueryParams: map[string]string{
			"search": name,
		},
		Body: nil,
	})

	res, err := pre.GetRespBody()
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &users)
	if err != nil {
		return
	}
	return
}

// 获取用户的群组
func (p *Gitlab) GetGroups() (groups []*models.GroupInfo, err error) {
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Method:  http.MethodGet,
		Url:     "/groups",
		Headers: nil,
		QueryParams: map[string]string{
			"order_by": "id",
		},
		Body: nil,
	})
	res, err := pre.GetRespBody()
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &groups)
	if err != nil {
		return
	}
	return
}

// 获取组下的项目
func (p *Gitlab) GetGroupProjects(groupID int) (projects []*models.ProjectInfo, err error) {
	url := fmt.Sprintf("/groups/%d/projects", groupID)
	opt := httpP.RequestOption{
		Method:  http.MethodGet,
		Url:     url,
		Headers: nil,
		QueryParams: map[string]string{
			"simple":            "true",
			"order_by":          "id",
			"include_subgroups": "true",
		},
		Body: nil,
	}
	pre := httpP.NewPreRequest(p.pre, opt)
	res, err := pre.GetRespBody()
	if err != nil {
		return
	}
	var resProjects []*models.ProjectInfo
	err = json.Unmarshal(res, &resProjects)
	if err != nil {
		return
	}
	for len(resProjects) != 0 {
		projects = append(projects, resProjects...)

		opt.QueryParams["id_before"] = strconv.Itoa(resProjects[len(resProjects)-1].ID)
		pre = httpP.NewPreRequest(p.pre, opt)
		var res []byte
		res, err = pre.GetRespBody()
		if err != nil {
			return
		}
		err = json.Unmarshal(res, &resProjects)
		if err != nil {
			return
		}
	}
	return
}

// 获取能够查看的项目
func (p *Gitlab) GetProjects() (projects []*models.ProjectInfo, err error) {
	url := "/projects"
	opt := httpP.RequestOption{
		Url:    url,
		Method: http.MethodGet,
		QueryParams: map[string]string{
			"simple":     "true",
			"membership": "true",
			"statistics": "true",
			"pagination": "keyset",
			"order_by":   "id",
			"sort":       "desc",
			"per_page":   "10",
		},
	}
	pre := httpP.NewPreRequest(p.pre, opt)
	res, err := pre.GetRespBody()
	if err != nil {
		return
	}
	var resProjects []*models.ProjectInfo
	err = json.Unmarshal(res, &resProjects)
	if err != nil {
		return
	}
	for len(resProjects) != 0 {
		projects = append(projects, resProjects...)

		opt.QueryParams["id_before"] = strconv.Itoa(resProjects[len(resProjects)-1].ID)
		pre = httpP.NewPreRequest(p.pre, opt)
		var res []byte
		res, err = pre.GetRespBody()
		if err != nil {
			return
		}
		err = json.Unmarshal(res, &resProjects)
		if err != nil {
			return
		}
	}
	return
}

// 创建合并请求
func (p *Gitlab) CreateMerge(body models.MergeRequest) (err error) {

	url := fmt.Sprintf("/projects/%d/merge_requests", body.Id)
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:     url,
		Method:  http.MethodPost,
		Headers: nil,
		Body:    body,
	})
	res, err := pre.GetRespBody()
	if err != nil {

		return
	}
	var resInfo models.MergeResInfo
	err = json.Unmarshal(res, &resInfo)
	if err != nil {
		return
	}

	if len(resInfo.Message) != 0 {
		if strings.Contains(fmt.Sprintf("%v", resInfo.Message), "already exists") {
			//TODO:请求已存在是否需要打印
			return
		}
		err = fmt.Errorf("创建合并请求失败:%v", resInfo.Message)
		return
	}
	return
}

// 获取分支详情
func (p *Gitlab) GetBranch(projectID int, branchName string) (branch *models.BranchInfo, err error) {
	url := fmt.Sprintf("/projects/%d/repository/branches/%s", projectID, branchName)
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:         url,
		Method:      http.MethodGet,
		Headers:     nil,
		QueryParams: nil,
		Body:        nil,
	})
	res, err := pre.GetRespBody()
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &branch)
	if err != nil {
		return
	}
	return
}

// 获取提交存在的分支
func (p *Gitlab) GetCommitBranches(projectID int, sha string) (branches []string, err error) {
	url := fmt.Sprintf("/projects/%d/repository/commits/%s/refs", projectID, sha)
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:     url,
		Method:  http.MethodGet,
		Headers: nil,
		QueryParams: map[string]string{
			"type": "branch", //类型 branch,tag,all 默认是all
		},
		Body: nil,
	})
	res, err := pre.GetRespBody()
	if err != nil {
		return
	}
	var resInfo []struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}
	err = json.Unmarshal(res, &resInfo)
	if err != nil {
		return
	}
	for _, v := range resInfo {
		branches = append(branches, v.Name)
	}
	return
}

// 自动提交合并请求
func (p *Gitlab) AutoMarge(req models.MergeRequest) (err error) {
	branchInfo, err := p.GetBranch(req.Id, req.SourceBranch)
	if err != nil {
		return err
	}

	mrBranchs, err := p.GetCommitBranches(req.Id, branchInfo.Commit.ID)
	if err != nil {
		return err
	}
	for _, branch := range mrBranchs {
		if branch == req.TargetBranch {
			// 目标分支没有落后
			return nil
		}
	}
	err = p.CreateMerge(req)
	if err != nil {
		return err
	}

	return
}
