package custommodel

import "gorm.io/gorm"

//============= CREATING  JOB TABLE ======================================
type Job struct {
	gorm.Model
	JobTitle  string  `json:"job_title" validate:"required"`
	JobSalary string  `json:"job_salary" validate:"required"`
	Company   Company `gorm:"ForeignKey:uid"`
	Uid       uint64  `JSON:"uid, omitempty"`
}

//================ CREATING COMAPNY TABLE =================================
type Company struct {
	gorm.Model
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

//=================== CREATE COMPANY FEIDLS ===================================
type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateJob struct {
	JobTitle  string `json:"job_title" validate:"required"`
	JobSalary string `json:"job_salary" validate:"required"`
}
