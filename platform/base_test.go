package platform

//func TestAuthMarge(t *testing.T) {
//	type args struct {
//		p   Base
//		req models.MergeRequest
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "test",
//			args: args{
//				p: NewGitlab(),
//				req: models.MergeRequest{
//					Id:                 437,
//					SourceBranch:       "dev",
//					TargetBranch:       "master",
//					Title:              "auto dev->master",
//					AssigneeId:         0,
//					ReviewerIds:        nil,
//					RemoveSourceBranch: false,
//				},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := AuthMarge(tt.args.p, tt.args.req); (err != nil) != tt.wantErr {
//				t.Errorf("AuthMarge() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
