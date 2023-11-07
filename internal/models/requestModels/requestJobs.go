package models

type NewJobs struct {
	Jobname         string               `json:"jobname" validate:"required"`
	MinNoticePeriod string               `json:"min_notice_period" validate:"required"`
	MaxNoticePeriod string               `json:"max_notice_period" validate:"required"`
	Location        []NewLocation        `json:"location" `
	TechnologyStack []NewTechnologyStack `json:"technology_stack" `
	Description     string               `json:"description" validate:"required"`
	MinExperience   string               `json:"min_experience" validate:"required"`
	MaxExperience   string               `json:"max_experience" validate:"required"`
	Qualifications  []NewQualification   `json:"qualifications"`
	Shift           []NewShift           `json:"shift"`
	Jobtype         string               `json:"jobtype" validate:"required"`
}

type NewLocation struct {
	PlaceId uint `json:"place_lid"`
}

type NewTechnologyStack struct {
	StackId uint `json:"stack_id"`
}

type NewQualification struct {
	QualificationId uint `json:"qualification_id"`
}

type NewShift struct {
	ShiftId uint `json:"shift_id"`
}
