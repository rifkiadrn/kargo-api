package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJobsReturnNotEmpty(t *testing.T) {
	repo := NewRepository()
	jobs, err := repo.GetJobs()

	assert.NotEmpty(t, jobs)
	assert.Nil(t, err)
}

func TestGetBidsReturnNotNil(t *testing.T) {
	repo := NewRepository()
	bids, err := repo.GetBids(1)

	assert.NotEmpty(t, bids)
	assert.Nil(t, err)
}
