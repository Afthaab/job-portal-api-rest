package service

import (
	"context"
	"errors"
	"sync"

	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	"github.com/rs/zerolog/log"
)

var cacheMap = make(map[uint]models.Jobs)

func (s *Service) ProccessApplication(ctx context.Context, applicationData []newModels.NewUserApplication) ([]newModels.NewUserApplication, error) {
	var wg = new(sync.WaitGroup)
	ch := make(chan newModels.NewUserApplication)
	var finalData []newModels.NewUserApplication

	for _, v := range applicationData {
		wg.Add(1)
		go func(v newModels.NewUserApplication) {
			defer wg.Done()

			val, exists := cacheMap[v.Jid]

			if !exists {
				jobData, err := s.UserRepo.GetTheJobData(v.Jid)
				if err != nil {
					return
				}
				cacheMap[v.Jid] = jobData
				val = jobData
			}
			check := compareAndCheck(v, val)
			if check {
				ch <- v
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		finalData = append(finalData, v)
	}

	return finalData, nil
}

func compareAndCheck(applicationData newModels.NewUserApplication, val models.Jobs) bool {
	if applicationData.Jobs.Experience < val.MinExperience {
		return false
	}
	if applicationData.Jobs.NoticePeriod < val.MinNoticePeriod {
		log.Info().Err(errors.New("experience did not match")).Send()
		return false
	}
	var count int
	count = compareLocations(applicationData.Jobs.Location, val.Location)
	if count == 0 {
		return false
	}

	count = compareQualifications(applicationData.Jobs.Qualifications, val.Qualifications)
	if count == 0 {
		return false
	}
	count = compareTechStack(applicationData.Jobs.TechnologyStack, val.TechnologyStack)
	if count == 0 {
		return false
	}
	count = compareShifts(applicationData.Jobs.Shift, val.Shift)
	if count == 0 {
		return false
	}

	return true
}

func compareLocations(locationsID []uint, val []models.Location) int {
	count := 0
	for _, v := range locationsID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

func compareQualifications(qualificationID []uint, val []models.Qualification) int {
	count := 0
	for _, v := range qualificationID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

func compareTechStack(stackID []uint, val []models.TechnologyStack) int {
	count := 0
	for _, v := range stackID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

func compareShifts(shiftID []uint, val []models.Shift) int {
	count := 0
	for _, v := range shiftID {
		for _, v1 := range val {
			if v == v1.ID {
				count++
			}
		}
	}
	return count
}

// check, v, err := s.compareAndCheck(v)

// if err != nil {
// 	return nil, err
// }
// if check {
// 	finalData = append(finalData, v)
// }
