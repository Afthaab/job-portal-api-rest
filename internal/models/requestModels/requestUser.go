package models

import (
	"gorm.io/gorm"
)

type NewJobs struct {
	Cid             uint                 `json:"cid"`
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
	gorm.Model
	PlaceName string `json:"place_name"`
}

type NewTechnologyStack struct {
	gorm.Model
	StackName string `json:"stack_name"`
}

type NewQualification struct {
	gorm.Model
	QualificationRequired string `json:"qualification_required"`
}

type NewShift struct {
	gorm.Model
	ShiftType string `json:"shift_type"`
}
