package service

import (
	"context"
	"sync"

	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
)

func (s *Service) ProccessApplication(ctx context.Context, applicationData []newModels.NewUserApplication) ([]newModels.NewUserApplication, error) {
	var wg = new(sync.WaitGroup)
	var finalData []newModels.NewUserApplication
	for _, v := range applicationData {
		wg.Add(1)
		go func(v newModels.NewUserApplication) {
			defer wg.Done()
			check, v, err := s.compareAndCheck(v)
			if err != nil {
				return
			}
			if check {
				finalData = append(finalData, v)
			}
		}(v)

		// check, v, err := s.compareAndCheck(v)

		// if err != nil {
		// 	return nil, err
		// }
		// if check {
		// 	finalData = append(finalData, v)
		// }
	}
	wg.Wait()
	return finalData, nil
}

var cacheMap = make(map[uint]models.Jobs)

func (s *Service) compareAndCheck(applicationData newModels.NewUserApplication) (bool, newModels.NewUserApplication, error) {
	val, exists := cacheMap[applicationData.Jid]
	if !exists {
		jobData, err := s.UserRepo.GetTheJobData(applicationData.Jid)
		if err != nil {
			return false, newModels.NewUserApplication{}, err
		}
		cacheMap[applicationData.Jid] = jobData
		val = jobData
	}
	if applicationData.Jobs.Experience < val.MinExperience {
		return false, newModels.NewUserApplication{}, nil
	}
	if applicationData.Jobs.Jobtype != val.Jobtype {
		return false, newModels.NewUserApplication{}, nil
	}
	if applicationData.Jobs.NoticePeriod < val.MinNoticePeriod {
		return false, newModels.NewUserApplication{}, nil
	}
	count := 0
	for _, v := range applicationData.Jobs.Location {
		for _, v1 := range val.Location {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, newModels.NewUserApplication{}, nil
	}
	count = 0
	for _, v := range applicationData.Jobs.Qualifications {
		for _, v1 := range val.Qualifications {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, newModels.NewUserApplication{}, nil
	}
	count = 0
	for _, v := range applicationData.Jobs.TechnologyStack {
		for _, v1 := range val.TechnologyStack {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, newModels.NewUserApplication{}, nil
	}
	count = 0
	for _, v := range applicationData.Jobs.Shift {
		for _, v1 := range val.Shift {
			if v == v1.ID {
				count++
			}
		}
	}
	if count == 0 {
		return false, newModels.NewUserApplication{}, nil
	}

	return true, applicationData, nil
}
