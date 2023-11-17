package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/cache/mockmodels"
	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	cachMock "github.com/afthaab/job-portal/internal/repository/mockModels"
	"github.com/redis/go-redis/v9"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestService_ProccessApplication(t *testing.T) {
	type args struct {
		ctx             context.Context
		applicationData []newModels.NewUserApplication
	}
	tests := []struct {
		name              string
		args              args
		want              []newModels.NewUserApplication
		mockRepoResponse  func() (models.Jobs, error)
		mockCacheResponse func() (string, error)
		wantErr           bool
	}{
		{
			name: "not found in redis and in database",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "afthab",
						Age:  "21",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:         "senior web developer",
							NoticePeriod:    7,
							Location:        []uint{1, 2},
							TechnologyStack: []uint{1, 2},
							Experience:      uint(5),
							Qualifications:  []uint{1, 2},
							Shift:           []uint{1, 2},
							Jobtype:         "permanent",
						},
					},
				},
			},
			want: nil,
			mockCacheResponse: func() (string, error) {
				return "", redis.Nil
			},
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("test error from cache mock")
			},
			wantErr: false,
		},
		{
			name: "not found in redis and in database",
			args: args{
				ctx: context.Background(),
				applicationData: []newModels.NewUserApplication{
					{
						Name: "afthab",
						Age:  "21",
						Jid:  1,
						Jobs: newModels.Requestfield{
							Jobname:         "senior web developer",
							NoticePeriod:    7,
							Location:        []uint{1, 2},
							TechnologyStack: []uint{1, 2},
							Experience:      uint(5),
							Qualifications:  []uint{1, 2},
							Shift:           []uint{1, 2},
							Jobtype:         "permanent",
						},
					},
				},
			},
			want: nil,
			mockCacheResponse: func() (string, error) {
				return "", redis.Nil
			},
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Model: gorm.Model{
						ID: 1,
					},
					Jobname:         "senior web developer",
					MinNoticePeriod: uint(5),
					MaxNoticePeriod: uint(10),
					Location: []models.Location{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					TechnologyStack: []models.TechnologyStack{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					MinExperience: uint(3),
					MaxExperience: uint(10),
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					Shift: []models.Shift{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
					},
					Jobtype: "permanent",
				}, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockCach := mockmodels.NewMockCaching(mc)
			mockCach.EXPECT().GetTheCacheData(gomock.Any(), gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
			mockCach.EXPECT().AddToTheCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("test error")).AnyTimes()

			mockRepo := cachMock.NewMockUserRepo(mc)
			mockRepo.EXPECT().GetTheJobData(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

			svc, err := NewService(mockRepo, &auth.Auth{}, mockCach)
			if err != nil {
				t.Error("could not initialize the mock repo layer")
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
