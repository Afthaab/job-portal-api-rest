package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/afthaab/job-portal/internal/auth"
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
			// [{Afthab 22 1 {web developer 10 [1 2] [1 2] 7 [1 2] [1 2] daily shift}}]
			want:    nil,
			wantErr: false,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("test error from mock function")
			},
		},
		{
			name: "success from mock",
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
					{
						Name: "Jeevan",
						Age:  "22",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:      "web developer",
							NoticePeriod: uint(0),
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
			want: []newModels.NewUserApplication{
				{
					Name: "Afthab",
					Age:  "22",
					Jid:  1,
					Jobs: newModels.Requestfield{
						Jobname:         "web developer",
						NoticePeriod:    uint(10),
						Location:        []uint{1, 2},
						TechnologyStack: []uint{1, 2},
						Experience:      7,
						Qualifications:  []uint{1, 2},
						Shift:           []uint{1, 2},
						Jobtype:         "daily shift",
					},
				},
				{
					Name: "Jeevan",
					Age:  "22",
					Jid:  1,
					Jobs: newModels.Requestfield{
						Jobname:         "web developer",
						NoticePeriod:    uint(10),
						Location:        []uint{1, 2},
						TechnologyStack: []uint{1, 2},
						Experience:      7,
						Qualifications:  []uint{1, 2},
						Shift:           []uint{1, 2},
						Jobtype:         "daily shift",
					},
				},
			},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := mockrepository.NewMockUserRepo(mc)

			mockRepo.EXPECT().GetTheJobData(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{})
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
