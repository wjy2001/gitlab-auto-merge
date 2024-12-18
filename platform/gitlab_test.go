package platform

import (
	"gitlab-auto-merge/conf"
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/pkg/structP"
	"reflect"
	"testing"
)

func init() {
	conf.InitConfig()
}
func TestGitlab_CreateMerge(t *testing.T) {

	type args struct {
		Body models.MergeRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				Body: models.MergeRequest{
					Id:                 437,
					SourceBranch:       "dev",
					TargetBranch:       "master",
					Title:              "test1",
					AssigneeId:         0,
					ReviewerIds:        nil,
					RemoveSourceBranch: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			user, _ := p.GetOwnInfo()
			tt.args.Body.AssigneeId = user.ID
			if err := p.CreateMerge(tt.args.Body); (err != nil) != tt.wantErr {
				t.Errorf("CreateMerge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitlab_GetBranch(t *testing.T) {
	type args struct {
		projectID  int
		branchName string
	}
	tests := []struct {
		name       string
		args       args
		wantBranch *models.BranchInfo
		wantErr    bool
	}{
		{
			name: "test",
			args: args{
				projectID:  182,
				branchName: "dev",
			},
			wantBranch: &models.BranchInfo{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotBranch, err := p.GetBranch(tt.args.projectID, tt.args.branchName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			structP.FmtStruct(gotBranch)
			if !reflect.DeepEqual(gotBranch, tt.wantBranch) {
				t.Errorf("GetBranch() gotBranch = %v, want %v", gotBranch, tt.wantBranch)
			}
		})
	}
}

func TestGitlab_GetCommitBranches(t *testing.T) {
	type args struct {
		projectID int
		sha       string
	}
	tests := []struct {
		name         string
		args         args
		wantBranches []string
		wantErr      bool
	}{
		{
			name: "test",
			args: args{
				projectID: 110,
				sha:       "279d963a71d4d4f054a5cea708268b7cd46a2279",
			},
			wantBranches: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotBranches, err := p.GetCommitBranches(tt.args.projectID, tt.args.sha)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotBranches {
				t.Log(v)
			}
			if !reflect.DeepEqual(gotBranches, tt.wantBranches) {
				t.Errorf("GetCommitBranches() gotBranches = %v, want %v", gotBranches, tt.wantBranches)
			}
		})
	}
}

func TestGitlab_GetGroupProjects(t *testing.T) {

	type args struct {
		groupID int
	}
	tests := []struct {
		name         string
		args         args
		wantProjects []*models.ProjectInfo
		wantErr      bool
	}{
		{
			name:         "bas",
			args:         args{groupID: 61},
			wantProjects: nil,
			wantErr:      false,
		},
		{
			name:         "easm",
			args:         args{groupID: 99},
			wantProjects: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotProjects, err := p.GetGroupProjects(tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotProjects {
				structP.FmtStruct(v)
			}
		})
	}
}

func TestGitlab_GetGroups(t *testing.T) {
	tests := []struct {
		name       string
		wantGroups []*models.GroupInfo
		wantErr    bool
	}{
		{
			name:       "test",
			wantGroups: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotGroups, err := p.GetGroups()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotGroups {
				t.Log(v)
				structP.FmtStruct(v)
			}
		})
	}
}

func TestGitlab_GetOwnInfo(t *testing.T) {
	tests := []struct {
		name     string
		wantUser *models.UserInfo
		wantErr  bool
	}{
		{
			name:     "test",
			wantUser: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotUser, err := p.GetOwnInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOwnInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			structP.FmtStruct(gotUser)
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetOwnInfo() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestGitlab_GetProjects(t *testing.T) {
	tests := []struct {
		name         string
		wantProjects []*models.ProjectInfo
		wantErr      bool
	}{
		{
			name:         "test",
			wantProjects: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotProjects, err := p.GetProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotProjects {
				t.Log(v)
				structP.FmtStruct(v)
			}
			if !reflect.DeepEqual(gotProjects, tt.wantProjects) {
				t.Errorf("GetProjects() gotProjects = %v, want %v", gotProjects, tt.wantProjects)
			}
		})
	}
}

func TestGitlab_GetUserByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantUsers []*models.UserInfo
		wantErr   bool
	}{
		{
			name:      "test",
			args:      args{name: "wang"},
			wantUsers: nil,
			wantErr:   false,
		},
		{
			name:      "test2",
			args:      args{name: "chen"},
			wantUsers: nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotUsers, err := p.GetUserByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, v := range gotUsers {
				t.Log(v)
				structP.FmtStruct(v)
			}
			if !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("GetUserByName() gotUsers = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
}

func TestGitlab_GetProjectBranches(t *testing.T) {
	type args struct {
		projectID int
	}
	tests := []struct {
		name         string
		args         args
		wantBranches []*models.BranchInfo
		wantErr      bool
	}{
		{
			name: "test",
			args: args{
				projectID: 177,
			},
			wantBranches: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotBranches, err := p.GetProjectBranches(tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBranches, tt.wantBranches) {
				t.Errorf("GetProjectBranches() gotBranches = %v, want %v", gotBranches, tt.wantBranches)
			}
		})
	}
}

func TestGitlab_GetGroupsMerges(t *testing.T) {
	type args struct {
		groupID int
		req     models.GetMergeReq
	}
	tests := []struct {
		name       string
		args       args
		wantMerges []models.MergeInfo
		wantErr    bool
	}{
		{
			name: "test uat",
			args: args{
				groupID: 61,
				req: models.GetMergeReq{
					Search:       "auto test to uat",
					State:        "opened",
					SourceBranch: "test",
					TargetBranch: "uat",
				},
			},
			wantMerges: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewGitlab()
			gotMerges, err := p.GetGroupsMerges(tt.args.groupID, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupsMerges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMerges, tt.wantMerges) {
				t.Errorf("GetGroupsMerges() gotMerges = %v, want %v", gotMerges, tt.wantMerges)
			}
		})
	}
}
