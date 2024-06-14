package models

type UserInfo struct {
	ID       int    `json:"id"`
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
