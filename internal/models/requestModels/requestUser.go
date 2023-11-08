package models

type NewUserApplication struct {
	Name string       `json:"name"`
	Age  string       `json:"age"`
	Jid  uint         `json:"jid"`
	Jobs Requestfield `json:"job_application"`
}

type Requestfield struct {
	Jobname         string `json:"jobName" validate:"required"`
	NoticePeriod    uint   `json:"noticePeriod" validate:"required"`
	Location        []uint `json:"location" `
	TechnologyStack []uint `json:"technologyStack" `
	Experience      uint   `json:"experience" validate:"required"`
	Qualifications  []uint `json:"qualifications"`
	Shift           []uint `json:"shifts"`
	Jobtype         string `json:"jobtype" validate:"required"`
}
