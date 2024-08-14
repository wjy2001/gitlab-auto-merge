package platform

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab-auto-merge/conf"
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/pkg/httpP"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	gitlabBasicAuth = "/api/v4"
)

type Gitlab struct {
	pre *resty.Client
}

func (p *Gitlab) NewPre() {
	c := conf.GetConfig()
	baseUrl := fmt.Sprint(c.Parameter.BasicUrl, gitlabBasicAuth)
	baseHeaders := map[string]string{
		"Private-Token": c.Parameter.Token,
	}
	p.pre = httpP.NewPreRequestClient(httpP.InitRequest{
		BaseURL:         baseUrl,
		BaseHeaders:     baseHeaders,
		BaseQueryParams: nil,
	})
	if c.Parameter.TLSSkiPVerify {
		log.Println("跳过证书验证")
		p.pre.SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}
	return
}
func NewGitlab() *Gitlab {
	gitlab := &Gitlab{}
	gitlab.NewPre()
	return gitlab
}

// GetOwnInfo 获取自己的信息
func (p *Gitlab) GetOwnInfo() (user models.UserInfo, err error) {
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Method: http.MethodGet,
		Url:    "/user",
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

// GetUserByName 获取用户的信息
func (p *Gitlab) GetUserByName(name string) (users []models.UserInfo, err error) {
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Method: http.MethodGet,
		Url:    "/users",
		QueryParams: map[string]string{
			"search": name,
		},
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

// GetGroups 获取用户的群组
func (p *Gitlab) GetGroups() (groups []models.GroupInfo, err error) {
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Method: http.MethodGet,
		Url:    "/groups",
		QueryParams: map[string]string{
			"order_by": "id",
		},
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

// GetGroupProjects 获取组下的项目
func (p *Gitlab) GetGroupProjects(groupID int) (projects []models.ProjectInfo, err error) {
	urlStr := fmt.Sprintf("/groups/%d/projects", groupID)
	opt := httpP.RequestOption{
		Method: http.MethodGet,
		Url:    urlStr,
		QueryParams: map[string]string{
			"simple":            "true",  //获取简单信息
			"order_by":          "id",    //按id排序
			"include_subgroups": "true",  //获取子组下的项目
			"archived":          "false", //获取没有归档的项目
			"per_page":          "20",    //每次获取20条
		},
	}
	pre := httpP.NewPreRequest(p.pre, opt)

	res := []byte{}
	resProjects := make([]models.ProjectInfo, 0, 20)
	for {
		if res, err = pre.GetRespBody(); err != nil {
			return
		}
		if err = json.Unmarshal(res, &resProjects); err != nil {
			return
		}
		if len(resProjects) == 0 {
			//如果没有项目了，跳出循环
			break
		}

		projects = append(projects, resProjects...)

		opt.QueryParams["id_before"] = strconv.Itoa(resProjects[len(resProjects)-1].ID)
		pre = httpP.NewPreRequest(p.pre, opt)
	}
	return
}

// GetProjects 获取能够查看的项目
func (p *Gitlab) GetProjects() (projects []models.ProjectInfo, err error) {
	urlStr := "/projects"
	opt := httpP.RequestOption{
		Url:    urlStr,
		Method: http.MethodGet,
		QueryParams: map[string]string{
			"simple":     "true",
			"membership": "true",
			"statistics": "true",
			"pagination": "keyset",
			"order_by":   "id",
			"sort":       "desc",
			"per_page":   "20",
		},
	}
	pre := httpP.NewPreRequest(p.pre, opt)

	res := []byte{}
	resProjects := make([]models.ProjectInfo, 0, 20)
	for {
		if res, err = pre.GetRespBody(); err != nil {
			return
		}
		if err = json.Unmarshal(res, &resProjects); err != nil {
			return
		}
		if len(resProjects) == 0 {
			//如果没有项目了，跳出循环
			break
		}

		projects = append(projects, resProjects...)

		opt.QueryParams["id_before"] = strconv.Itoa(resProjects[len(resProjects)-1].ID)
		pre = httpP.NewPreRequest(p.pre, opt)
	}
	return
}

// CreateMerge 创建合并请求
func (p *Gitlab) CreateMerge(body models.MergeRequest) (err error) {

	urlStr := fmt.Sprintf("/projects/%d/merge_requests", body.Id)
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:    urlStr,
		Method: http.MethodPost,
		Body:   body,
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

// GetBranch 获取分支详情
func (p *Gitlab) GetBranch(projectID int, branchName string) (branch models.BranchInfo, err error) {
	branchName = url.PathEscape(branchName)
	urlStr := fmt.Sprintf("/projects/%d/repository/branches/%s", projectID, branchName)
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:    urlStr,
		Method: http.MethodGet,
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

// GetProjectBranches 获取项目分支
func (p *Gitlab) GetProjectBranches(projectID int) (branches []models.BranchInfo, err error) {
	urlStr := fmt.Sprintf("/projects/%d/repository/branches", projectID)
	var page int64 = 1
	var pageSize int64 = 20
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:    urlStr,
		Method: http.MethodGet,
		QueryParams: map[string]string{
			"per_page": strconv.FormatInt(pageSize, 10), //每次获取20条
			"page":     strconv.FormatInt(page, 10),     //页数
		},
	})
	var res []byte
	resInfo := make([]models.BranchInfo, 0, pageSize)
	for {
		if res, err = pre.GetRespBody(); err != nil {
			return
		}
		if err = json.Unmarshal(res, &resInfo); err != nil {
			return
		}
		if len(resInfo) == 0 {
			//如果没有项目了，跳出循环
			break
		}
		branches = append(branches, resInfo...)
		page = page + 1
		pre.Option.QueryParams["page"] = strconv.FormatInt(page, 10)
	}

	return
}

// GetCommitBranches 获取提交存在的分支
func (p *Gitlab) GetCommitBranches(projectID int, sha string) (branches []string, err error) {
	urlStr := fmt.Sprintf("/projects/%d/repository/commits/%s/refs", projectID, sha)
	pre := httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:    urlStr,
		Method: http.MethodGet,
		QueryParams: map[string]string{
			"type": "branch", //类型 branch,tag,all 默认是all
		},
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

// AutoMerge 自动提交合并请求
func (p *Gitlab) AutoMerge(req models.MergeRequest) (err error) {
	var sourceBranchInfo models.BranchInfo
	var targetBranchInfo models.BranchInfo
	branches, err := p.GetProjectBranches(req.Id)
	if err != nil {
		return err
	}
	for _, branch := range branches {
		if branch.Name == req.SourceBranch {
			sourceBranchInfo = branch
		}
		if branch.Name == req.TargetBranch {
			targetBranchInfo = branch
		}
	}
	if len(sourceBranchInfo.Name) == 0 || len(targetBranchInfo.Name) == 0 {
		// 有分支不存在
		return nil
	}
	mrBranches, err := p.GetCommitBranches(req.Id, sourceBranchInfo.Commit.ID)
	if err != nil {
		return err
	}
	for _, branch := range mrBranches {
		if branch == req.TargetBranch {
			// 目标分支没有落后
			return nil
		}
	}
	//当指定title为空时，使用最新的title
	if len(req.Title) == 0 {
		req.Title = sourceBranchInfo.Commit.Title
	}
	err = p.CreateMerge(req)
	if err != nil {
		return err
	}

	return
}

// GetGroupsMerges 获取群组合并请求
func (p *Gitlab) GetGroupsMerges(groupID int, req models.GetMergeReq) (merges []models.MergeInfo, err error) {

	urlStr := fmt.Sprintf("/groups/%d/merge_requests", groupID)
	var page int64 = 1
	var pageSize int64 = 20
	reqMap := req.ToStringMap()
	reqMap["per_page"] = strconv.FormatInt(pageSize, 10)
	reqMap["page"] = strconv.FormatInt(page, 10)

	opt := httpP.RequestOption{
		Url:         urlStr,
		Method:      http.MethodGet,
		QueryParams: reqMap,
	}
	pre := httpP.NewPreRequest(p.pre, opt)
	var res []byte
	resInfo := make([]models.MergeInfo, 0, pageSize)
	for {
		if res, err = pre.GetRespBody(); err != nil {
			return
		}
		if err = json.Unmarshal(res, &resInfo); err != nil {
			return
		}
		if len(resInfo) == 0 {
			//如果没有项目了，跳出循环
			break
		}
		merges = append(merges, resInfo...)
		page = page + 1
		pre.Option.QueryParams["page"] = strconv.FormatInt(page, 10)
	}
	return
}

// UpdateMergeRequest 修改合并请求
func (p *Gitlab) UpdateMergeRequest(projectID int, mergeID int, req models.MergeRequest) (err error) {
	urlStr := fmt.Sprintf("/projects/%d/merge_requests/%d", projectID, mergeID)
	_ = httpP.NewPreRequest(p.pre, httpP.RequestOption{
		Url:    urlStr,
		Method: http.MethodPut,
		Body:   req,
	})
	return
}
