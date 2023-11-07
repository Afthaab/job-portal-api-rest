package models

type NewUserApplication struct {
	Name string  `json:"name"`
	Age  string  `json:"age"`
	Jid  uint    `json:"jid"`
	Jobs NewJobs `json:"job_application"`
}
