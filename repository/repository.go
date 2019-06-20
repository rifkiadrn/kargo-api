package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kargo-api/models"
)

const (
	JobsJSON = "jobs.json"
	BidsJSON = "bids.json"
)

type Repository interface {
	GetJobs() ([]models.Job, error)
	GetBids(jobID int64) ([]models.Bid, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetJobs() ([]models.Job, error) {
	var jobs []models.Job

	jobsFileData, err := os.Open(JobsJSON)
	if err != nil {
		return jobs, err
	}

	jobsJSON, _ := ioutil.ReadAll(jobsFileData)
	_ = json.Unmarshal(jobsJSON, &jobs)
	return jobs, nil
}

func (repo *repository) GetBids(jobID int64) ([]models.Bid, error) {
	var bids []models.Bid

	bidsFileData, err := os.Open(BidsJSON)
	if err != nil {
		return bids, err
	}

	bidsJSON, _ := ioutil.ReadAll(bidsFileData)
	_ = json.Unmarshal(bidsJSON, &bids)
	return bids, nil
}
