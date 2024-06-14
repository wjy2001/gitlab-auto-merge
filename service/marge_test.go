package service

import (
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/platform"
	"testing"
	"time"
)

func TestService_CreateAutoMargeTask(t *testing.T) {

	type args struct {
		taskInfo *models.TaskAutoMarge
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				taskInfo: &models.TaskAutoMarge{
					ProjectIDs:   []int{437},
					GroupIDs:     nil,
					SourceBranch: "dev",
					TargetBranch: "master",
					Title:        "auto dev",
					ReviewerID:   nil,
					IntervalTime: time.Second * 2,
					CreatedTime:  time.Now(),
					Enable:       false,
					Cancel:       nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Service{
				platform: platform.NewGitlab(),
			}
			if err := p.CreateAutoMargeTask(tt.args.taskInfo); (err != nil) != tt.wantErr {
				t.Errorf("CreateAutoMargeTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		select {}
	}
}
