package processor

import (
	"fmt"

	"github.com/kargo-api/models"
	"github.com/kargo-api/repository"
)

type Processor interface {
	GetSortedJobs() ([]models.Job, error)
	GetSortedBids(jobID int64) ([]models.Bid, error)
}

type processor struct {
	repo repository.Repository
}

func NewProcessor() Processor {
	return &processor{
		repo: repository.NewRepository(),
	}
}

func (proc *processor) GetSortedJobs() ([]models.Job, error) {
	var sortedJobs []models.Job

	unsortedJob, err := proc.repo.GetJobs()
	if err != nil {
		return sortedJobs, err
	}
	quickSortJobs(&unsortedJob, 0, len(unsortedJob)-1)
	for index := 0; index < len(unsortedJob); index++ {
		fmt.Println(unsortedJob[index])
	}
	return unsortedJob, nil
}

func (proc *processor) GetSortedBids(jobID int64) ([]models.Bid, error) {
	var sortedBids []models.Bid

	unsortedBid, err := proc.repo.GetBids(jobID)
	if err != nil {
		return sortedBids, err
	}
	quickSortBids(&unsortedBid, 0, len(unsortedBid)-1)
	for index := 0; index < len(unsortedBid); index++ {
		fmt.Println(unsortedBid[index])
	}
	return unsortedBid, nil
}

func quickSortJobs(unsortedJob *[]models.Job, p int, r int) []models.Job {
	var sortedJobs []models.Job
	// for index := 0; index < len(unsortedJob); index++ {
	// 	fmt.Println(unsortedJob[index])
	// }
	if p < r {
		pivotIdx := partition(unsortedJob, p, r)
		quickSortJobs(unsortedJob, p, pivotIdx-1)
		quickSortJobs(unsortedJob, pivotIdx+1, r)
	}
	return sortedJobs
}

func partition(jobs *[]models.Job, low, high int) int {
	pivot := (*jobs)[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (*jobs)[j].ShipmentDate.Before(pivot.ShipmentDate) {
			i++
			jobI := (*jobs)[i]
			(*jobs)[i] = (*jobs)[j]
			(*jobs)[j] = jobI
		}
	}
	(*jobs)[high] = (*jobs)[i+1]
	(*jobs)[i+1] = pivot
	return i + 1
}

func quickSortBids(unsortedBids *[]models.Bid, p int, r int) []models.Bid {
	var sortedBids []models.Bid
	// for index := 0; index < len(unsortedJob); index++ {
	// 	fmt.Println(unsortedJob[index])
	// }
	if p < r {
		pivotIdx := partitionBid(unsortedBids, p, r)
		quickSortBids(unsortedBids, p, pivotIdx-1)
		quickSortBids(unsortedBids, pivotIdx+1, r)
	}
	return sortedBids
}

func partitionBid(bids *[]models.Bid, low, high int) int {
	pivot := (*bids)[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (*bids)[j].Price < pivot.Price {
			i++
			bidI := (*bids)[i]
			(*bids)[i] = (*bids)[j]
			(*bids)[j] = bidI
		}
	}
	(*bids)[high] = (*bids)[i+1]
	(*bids)[i+1] = pivot
	return i + 1
}
