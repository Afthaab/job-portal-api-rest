package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/cache"
	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	mockrepository "github.com/afthaab/job-portal/internal/repository/mockModels"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_ProccessApplication(t *testing.T) {
	type args struct {
		ctx             context.Context
		applicationData []newModels.NewUserApplication
	}
	tests := []struct {
		name             string
		args             args
		want             []newModels.NewUserApplication
		wantErr          bool
		mockRepoResponse func() (models.Jobs, error)
	}{
		{
			name: "error from mock function",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "Afthab",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(10),
							Location: []uint{
								uint(1), uint(2),
							},
							TechnologyStack: []uint{
								uint(1), uint(2),
							},
							Experience: uint(10),
							Qualifications: []uint{
								uint(1), uint(2),
							},
							Shift: []uint{
								uint(1), uint(2),
							},
							Jobtype: "daily shift",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("test error from mock function")
			},
		},
		{
			name: "failure in checking the job experience",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "Afthab",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(10),
							Location: []uint{
								uint(1), uint(2),
							},
							TechnologyStack: []uint{
								uint(1), uint(2),
							},
							Experience: uint(3),
							Qualifications: []uint{
								uint(1), uint(2),
							},
							Shift: []uint{
								uint(1), uint(2),
							},
							Jobtype: "daily shift",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:             1,
					Jobname:         "web developer",
					MinNoticePeriod: uint(5),
					MaxNoticePeriod: uint(20),
					Location: []models.Location{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					TechnologyStack: []models.TechnologyStack{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Description:   "test description",
					MinExperience: 5,
					MaxExperience: 15,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Jobtype: "test job type",
				}, nil
			},
		},
		{
			name: "failure in checking the notice period",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "Afthab",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(1),
							Location: []uint{
								uint(1), uint(2),
							},
							TechnologyStack: []uint{
								uint(1), uint(2),
							},
							Experience: uint(7),
							Qualifications: []uint{
								uint(1), uint(2),
							},
							Shift: []uint{
								uint(1), uint(2),
							},
							Jobtype: "daily shift",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:             1,
					Jobname:         "web developer",
					MinNoticePeriod: uint(5),
					MaxNoticePeriod: uint(20),
					Location: []models.Location{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					TechnologyStack: []models.TechnologyStack{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Description:   "test description",
					MinExperience: 5,
					MaxExperience: 15,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Jobtype: "test job type",
				}, nil
			},
		},
		{
			name: "failure in checking the locations",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "Afthab",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(6),
							Location:     []uint{},
							TechnologyStack: []uint{
								uint(1), uint(2),
							},
							Experience: uint(7),
							Qualifications: []uint{
								uint(1), uint(2),
							},
							Shift: []uint{
								uint(1), uint(2),
							},
							Jobtype: "daily shift",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:             1,
					Jobname:         "web developer",
					MinNoticePeriod: uint(5),
					MaxNoticePeriod: uint(20),
					Location: []models.Location{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					TechnologyStack: []models.TechnologyStack{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Description:   "test description",
					MinExperience: 5,
					MaxExperience: 15,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Jobtype: "test job type",
				}, nil
			},
		},
		{
			name: "failure in checking the qualifications",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "Afthab",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(6),
							Location: []uint{
								uint(1), uint(2),
							},
							TechnologyStack: []uint{
								uint(1), uint(2),
							},
							Experience:     uint(7),
							Qualifications: []uint{},
							Shift: []uint{
								uint(1), uint(2),
							},
							Jobtype: "daily shift",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:             1,
					Jobname:         "web developer",
					MinNoticePeriod: uint(5),
					MaxNoticePeriod: uint(20),
					Location: []models.Location{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					TechnologyStack: []models.TechnologyStack{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Description:   "test description",
					MinExperience: 5,
					MaxExperience: 15,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Jobtype: "test job type",
				}, nil
			},
		},
		{
			name: "failure in checking the technologies",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "Afthab",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(6),
							Location: []uint{
								uint(1), uint(2),
							},
							TechnologyStack: []uint{},
							Experience:      uint(7),
							Qualifications: []uint{
								uint(1), uint(2),
							},
							Shift: []uint{
								uint(1), uint(2),
							},
							Jobtype: "daily shift",
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Company: models.Company{
						Model: gorm.Model{
							ID: 1,
						},
					},
					Cid:             1,
					Jobname:         "web developer",
					MinNoticePeriod: uint(5),
					MaxNoticePeriod: uint(20),
					Location: []models.Location{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					TechnologyStack: []models.TechnologyStack{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Description:   "test description",
					MinExperience: 5,
					MaxExperience: 15,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Jobtype: "test job type",
				}, nil
			},
		},
		// {
		// 	name: "success from mock",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		applicationData: []newModels.NewUserApplication{
		// 			{
		// 				Name: "Afthab",
		// 				Age:  "22",
		// 				Jid:  1,
		// 				Jobs: newModels.Requestfield{
		// 					Jobname:      "web developer",
		// 					NoticePeriod: uint(10),
		// 					Location: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					TechnologyStack: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					Experience: uint(7),
		// 					Qualifications: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					Shift: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					Jobtype: "daily shift",
		// 				},
		// 			},
		// 			{
		// 				Name: "Jeevan",
		// 				Age:  "22",
		// 				Jid:  2,
		// 				Jobs: newModels.Requestfield{
		// 					Jobname:      "web developer",
		// 					NoticePeriod: uint(10),
		// 					Location: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					TechnologyStack: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					Experience: uint(7),
		// 					Qualifications: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					Shift: []uint{
		// 						uint(1), uint(2),
		// 					},
		// 					Jobtype: "daily shift",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: []newModels.NewUserApplication{
		// 		{
		// 			Name: "Afthab",
		// 			Age:  "22",
		// 			Jid:  1,
		// 			Jobs: newModels.Requestfield{
		// 				Jobname:         "web developer",
		// 				NoticePeriod:    uint(10),
		// 				Location:        []uint{1, 2},
		// 				TechnologyStack: []uint{1, 2},
		// 				Experience:      7,
		// 				Qualifications:  []uint{1, 2},
		// 				Shift:           []uint{1, 2},
		// 				Jobtype:         "daily shift",
		// 			},
		// 		},
		// 		{
		// 			Name: "Jeevan",
		// 			Age:  "22",
		// 			Jid:  2,
		// 			Jobs: newModels.Requestfield{
		// 				Jobname:         "web developer",
		// 				NoticePeriod:    uint(10),
		// 				Location:        []uint{1, 2},
		// 				TechnologyStack: []uint{1, 2},
		// 				Experience:      7,
		// 				Qualifications:  []uint{1, 2},
		// 				Shift:           []uint{1, 2},
		// 				Jobtype:         "daily shift",
		// 			},
		// 		},
		// 	},
		// 	wantErr:          false,
		// 	mockRepoResponse: nil,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := mockrepository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetTheJobData(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			for _, v := range tt.args.applicationData {
				if v.Jid == uint(1) {
					mockRepo.EXPECT().GetTheJobData(v.Jid).Return(models.Jobs{
						Model: gorm.Model{
							ID: 1,
						},
						Company: models.Company{
							Model: gorm.Model{
								ID: 1,
							},
						},
						Cid:             1,
						Jobname:         "web developer",
						MinNoticePeriod: uint(5),
						MaxNoticePeriod: uint(20),
						Location: []models.Location{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						TechnologyStack: []models.TechnologyStack{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						Description:   "test description",
						MinExperience: 5,
						MaxExperience: 15,
						Qualifications: []models.Qualification{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						Shift: []models.Shift{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						Jobtype: "test job type",
					}, nil).AnyTimes()
				}
				if v.Jid == uint(2) {
					mockRepo.EXPECT().GetTheJobData(v.Jid).Return(models.Jobs{
						Model: gorm.Model{
							ID: 1,
						},
						Company: models.Company{
							Model: gorm.Model{
								ID: 1,
							},
						},
						Cid:             1,
						Jobname:         "web developer",
						MinNoticePeriod: uint(5),
						MaxNoticePeriod: uint(20),
						Location: []models.Location{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						TechnologyStack: []models.TechnologyStack{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						Description:   "test description",
						MinExperience: 5,
						MaxExperience: 15,
						Qualifications: []models.Qualification{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						Shift: []models.Shift{
							{
								Model: gorm.Model{
									ID: 1,
								},
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
							},
						},
						Jobtype: "test job type",
					}, nil)
				}
			}

			svc, err := NewService(mockRepo, &auth.Auth{}, &cache.RDBLayer{})
			if err != nil {
				t.Errorf("error is initializing the repo layer")
				return
			}
			got, err := svc.ProccessApplication(tt.args.ctx, tt.args.applicationData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ProccessApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ProccessApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
