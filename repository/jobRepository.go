package repository

import custommodel "project-gql/models"

//===================== COMPANY INTERFACE ==================================
type Company interface {
	CreateCompany(custommodel.Company) (custommodel.Company, error)
	GetAllCompany() ([]custommodel.Company, error)
	GetCompany(id int) (custommodel.Company, error)
	CreateJob(j custommodel.Job) (custommodel.Job, error)
	GetJobs(id int) ([]custommodel.Job, error)
	GetAllJobs() ([]custommodel.Job, error)
}

//=============== CREAT COMPANY FUNC IS USED TO CREATE COMPANY ACCOUNT INTHE DATABASE ========================
func (r *Repo) CreateCompany(u custommodel.Company) (custommodel.Company, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return custommodel.Company{}, err
	}
	return u, nil
}

//================= GET ALL COMPANY FUNC IS USDE TO RETRIVE ALL COMPANY IN THE DATABASE ======================
func (r *Repo) GetAllCompany() ([]custommodel.Company, error) {
	var s []custommodel.Company
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}

//=================== GET COMPANY FUNC IS USED TO GET COMPANY BY ID IN THE DATABASE =============================
func (r *Repo) GetCompany(id int) (custommodel.Company, error) {
	var m custommodel.Company

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return custommodel.Company{}, err
	}
	return m, nil

}

//==================== CREATE JOB FUNC IS USDE TO CREATE JOB IN THE COMPANY ==============================
func (r *Repo) CreateJob(j custommodel.Job) (custommodel.Job, error) {
	err := r.db.Create(&j).Error
	if err != nil {
		return custommodel.Job{}, err
	}
	return j, nil
}

//========================== GET JOBS FUNC IS USED TO RETRIVE THE JOB IN THE COMPANY ===============
func (r *Repo) GetJobs(id int) ([]custommodel.Job, error) {
	var m []custommodel.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil

}

//=============== GET ALL JOBS IS USED TO RETRIVE THE ALL JOBS THE DATABASE ========================
func (r *Repo) GetAllJobs() ([]custommodel.Job, error) {
	var s []custommodel.Job
	err := r.db.Find(&s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}
