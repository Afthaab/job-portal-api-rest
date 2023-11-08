package models

type NewJobs struct {
	Jobname         string `json:"jobName" validate:"required"`
	MinNoticePeriod uint   `json:"minNoticePeriod" validate:"required"`
	MaxNoticePeriod uint   `json:"maxNoticePeriod" validate:"required"`
	Location        []uint `json:"location" `
	TechnologyStack []uint `json:"technologyStack" `
	Description     string `json:"description" validate:"required"`
	MinExperience   uint   `json:"minExperience" validate:"required"`
	MaxExperience   uint   `json:"maxExperience" validate:"required"`
	Qualifications  []uint `json:"qualifications"`
	Shift           []uint `json:"shifts"`
	Jobtype         string `json:"jobtype" validate:"required"`
}

type ResponseNewJobs struct {
	Jobid uint `json:"jobID"`
}

// type NewLocation struct {
// 	PlaceId uint `json:"place_lid"`
// }

// type NewTechnologyStack struct {
// 	StackId uint `json:"stack_id"`
// }

// type NewQualification struct {
// 	QualificationId uint `json:"qualification_id"`
// }

// type NewShift struct {
// 	ShiftId uint `json:"shift_id"`
// }
